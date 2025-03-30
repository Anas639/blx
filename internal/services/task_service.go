package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/anas639/blx/internal/task"
)

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (this *TaskService) CreateTask(name string) (*task.Task, error) {
	stmt, err := this.db.Prepare("INSERT INTO tasks(name) VALUES(?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name)
	if err != nil {
		return nil, err
	}
	taskID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	t := task.NewTask(taskID, name)
	return t, nil
}

func (this *TaskService) DeleteTask(taskId int64) error {
	stmt, err := this.db.Prepare("DELETE FROM tasks WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(taskId)
	if err != nil {
		return err
	}
	if rcount, _ := res.RowsAffected(); rcount < 1 {
		return fmt.Errorf("No task found with id %d", taskId)
	}

	return nil
}

func (this *TaskService) GetTasks(filter task.TaskFilter) ([]*task.Task, error) {
	tasks := []*task.Task{}

	var statusFilter string
	if filter.Statuses != nil && len(filter.Statuses) > 0 {
		sfBuilder := strings.Builder{}
		sfBuilder.WriteString("where t.status in (")
		for i, status := range filter.Statuses {
			sfBuilder.WriteString(fmt.Sprintf("\"%s\"", status))
			if i < len(filter.Statuses)-1 {
				sfBuilder.WriteString(",")
			}
		}
		sfBuilder.WriteString(")")
		statusFilter = sfBuilder.String()
	}
	res, err := this.db.Query(fmt.Sprintf("select t.id,t.name,t.status, t.project_id, p.name as \"project_name\" from tasks t left join projects p on t.project_id = p.id %s", statusFilter))
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var id int64
		var name string
		var status string
		var project_id sql.NullInt64
		var project_name sql.NullString
		err = res.Scan(&id, &name, &status, &project_id, &project_name)

		if err != nil {
			return nil, err
		}

		task := task.NewTask(id, name)
		task.SetStatus(status)
		sessions, err := this.getTaskSessions(task.Id)
		if err == nil {
			task.SetSessions(sessions)
		}

		if project_id.Valid {
			task.SetProject(project_id.Int64, project_name.String)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (this *TaskService) StartTask(taskId int64) (*task.Task, error) {
	task, err := this.GetTaskById(taskId, task.TaskFilter{})
	if err != nil {
		return nil, err
	}
	session, err := task.Start()
	if err != nil {
		return nil, err
	}
	err = this.createTaskSession(session)
	if err != nil {
		return nil, err
	}
	this.updateTaskStatus(task)
	return task, nil
}

func (this *TaskService) EndTask(taskId int64) (*task.Task, error) {
	task, err := this.GetTaskById(taskId, task.TaskFilter{})
	if err != nil {
		return nil, err
	}

	sessions, err := this.getTaskSessions(taskId)
	if err != nil {
		return nil, err
	}
	task.SetSessions(sessions)

	session, err := task.End()
	if err != nil {
		return nil, err
	}
	err = this.endTaskSession(session)
	if err != nil {
		return nil, err
	}

	err = this.updateTaskStatus(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (this *TaskService) PauseTask(taskId int64) (*task.Task, error) {
	task, err := this.GetTaskById(taskId, task.TaskFilter{})
	if err != nil {
		return nil, err
	}

	sessions, err := this.getTaskSessions(taskId)
	if err != nil {
		return nil, err
	}
	task.SetSessions(sessions)

	session, err := task.Pause()
	if err != nil {
		return nil, err
	}

	err = this.endTaskSession(session)
	if err != nil {
		return nil, err
	}

	err = this.updateTaskStatus(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (this *TaskService) UpdateTaskName(taskId int64, name string) error {
	stmt, err := this.db.Prepare("UPDATE tasks SET name = ? where id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(name, taskId)
	if rcount, _ := res.RowsAffected(); rcount < 1 {
		return fmt.Errorf("No task found with id %d", taskId)
	}
	return nil
}

func (this *TaskService) AssignProject(taskId int64, projectId int64) error {
	stmt, err := this.db.Prepare("UPDATE tasks SET project_id = ? where id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(projectId, taskId)
	if rcount, _ := res.RowsAffected(); rcount < 1 {
		return fmt.Errorf("No task found with id %d", taskId)
	}
	return nil
}

func (this *TaskService) GetLastActiveTask() (*task.Task, error) {
	q := "select t.id,t.name,t.status, t.project_id, p.name as \"project_name\" from tasks t left join projects p on t.project_id = p.id inner join task_sessions ts on ts.task_id = t.id where t.status=\"ongoing\" and ts.start_time = (select MAX(start_time) from task_sessions where end_time is null);"
	row := this.db.QueryRow(q)
	var id int64
	var name string
	var status string
	var project_id sql.NullInt64
	var project_name sql.NullString
	err := row.Scan(&id, &name, &status, &project_id, &project_name)
	if err != nil {
		return nil, fmt.Errorf("No active task found")
	}
	t := task.NewTask(id, name)
	if project_id.Valid {
		t.SetProject(project_id.Int64, project_name.String)
	}
	t.SetStatus(status)
	sessions, err := this.getTaskSessions(t.Id)
	if err == nil {
		t.SetSessions(sessions)
	}
	return t, nil
}

func (this *TaskService) GetTaskById(taskId int64, filter task.TaskFilter) (*task.Task, error) {

	var statusFilter string

	if filter.Statuses != nil && len(filter.Statuses) > 0 {
		sfBuilder := strings.Builder{}
		sfBuilder.WriteString("and t.status in (")
		for i, status := range filter.Statuses {
			sfBuilder.WriteString(fmt.Sprintf("\"%s\"", status))
			if i < len(filter.Statuses)-1 {
				sfBuilder.WriteString(",")
			}
		}
		sfBuilder.WriteString(")")
		statusFilter = sfBuilder.String()
	}

	q := "select t.id,t.name,t.status, t.project_id, p.name as \"project_name\" from tasks t left join projects p on t.project_id = p.id where t.id = ? %s"
	stmt, err := this.db.Prepare(fmt.Sprintf(q, statusFilter))
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(taskId)
	var id int64
	var name string
	var status string
	var project_id sql.NullInt64
	var project_name sql.NullString
	err = row.Scan(&id, &name, &status, &project_id, &project_name)
	if err != nil {
		return nil, fmt.Errorf("No task found with id %d", taskId)
	}
	t := task.NewTask(id, name)
	if project_id.Valid {
		t.SetProject(project_id.Int64, project_name.String)
	}
	t.SetStatus(status)
	sessions, err := this.getTaskSessions(taskId)
	if err == nil {
		t.SetSessions(sessions)
	}
	return t, nil
}
func (this *TaskService) updateTaskStatus(task *task.Task) error {
	stmt, err := this.db.Prepare("UPDATE tasks SET status = ? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Status, task.Id)
	if err != nil {
		return err
	}
	return nil
}

func (this *TaskService) createTaskSession(session *task.TaskSession) error {
	stmt, err := this.db.Prepare("INSERT INTO task_sessions(task_id,start_time) VALUES(?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.TaskId, session.StartsAt)
	if err != nil {
		return err
	}
	return nil
}

func (this *TaskService) getTaskSessions(taskId int64) ([]*task.TaskSession, error) {
	stmt, err := this.db.Prepare("SELECT id,start_time,end_time,task_id FROM task_sessions where task_id = ?")
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	res, err := stmt.Query(taskId)
	if err != nil {
		return nil, err
	}

	sessions := []*task.TaskSession{}
	for res.Next() {
		var id, taskId int64
		var startTime, endTime sql.NullTime
		err = res.Scan(&id, &startTime, &endTime, &taskId)
		if err != nil {
			return nil, err
		}
		session := task.NewTaskSession(id, startTime.Time, endTime.Time, taskId)
		sessions = append(sessions, session)
	}
	return sessions, nil

}

func (this *TaskService) endTaskSession(session *task.TaskSession) error {
	stmt, err := this.db.Prepare("UPDATE task_sessions set end_time=? where id=?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.EndsAt, session.Id)
	return nil
}

func (this *TaskService) Search(keyword string, filter task.TaskFilter) ([]*task.Task, error) {
	tasks := []*task.Task{}

	var statusFilter string
	if filter.Statuses != nil && len(filter.Statuses) > 0 {
		sfBuilder := strings.Builder{}
		sfBuilder.WriteString("AND t.status in (")
		for i, status := range filter.Statuses {
			sfBuilder.WriteString(fmt.Sprintf("\"%s\"", status))
			if i < len(filter.Statuses)-1 {
				sfBuilder.WriteString(",")
			}
		}
		sfBuilder.WriteString(")")
		statusFilter = sfBuilder.String()
	}
	q := `
SELECT t.id,t.name,t.status, t.project_id, p.name as "project_name"
FROM tasks t
LEFT JOIN projects p ON t.project_id = p.id
WHERE (t.id IN (SELECT rowid FROM tasks_fts WHERE name MATCH '%s*') %s)
OR (p.id IS NOT NULL AND p.id IN (SELECT rowid FROM projects_fts WHERE name MATCH '%s*'));
	`
	res, err := this.db.Query(fmt.Sprintf(q, keyword, statusFilter, keyword))
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var id int64
		var name string
		var status string
		var project_id sql.NullInt64
		var project_name sql.NullString
		err = res.Scan(&id, &name, &status, &project_id, &project_name)

		if err != nil {
			return nil, err
		}

		task := task.NewTask(id, name)
		task.SetStatus(status)
		sessions, err := this.getTaskSessions(task.Id)
		if err == nil {
			task.SetSessions(sessions)
		}

		if project_id.Valid {
			task.SetProject(project_id.Int64, project_name.String)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
