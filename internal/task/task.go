package task

import (
	"fmt"
	"time"

	"github.com/anas639/blx/internal/project"
)

type Task struct {
	Id             int64
	Name           string
	Status         TaskStatus
	CreatedAt      time.Time
	sessions       []*TaskSession
	currentSession *TaskSession
	Project        *project.Project
}

func NewTask(id int64, name string) *Task {
	return &Task{
		Id:        id,
		Status:    TASK_NEW,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (this *Task) String() string {
	return fmt.Sprintf("%s [%s]", this.Name, this.Status)
}
func (this *Task) getLastSession() *TaskSession {
	if this.currentSession != nil {
		return this.currentSession
	}

	l := len(this.sessions)
	if l == 0 {
		return nil
	}
	return this.sessions[l-1]
}

func (this *Task) canStart() bool {
	return this.Status == TASK_NEW || this.Status == TASK_PAUSED
}

func (this *Task) canPause() bool {
	return this.Status == TASK_ONGOING
}

func (this *Task) canEnd() bool {
	return this.Status == TASK_PAUSED || this.Status == TASK_ONGOING
}

func (this *Task) Start() (*TaskSession, error) {
	if !this.canStart() {
		return nil, fmt.Errorf("This task can't be started")
	}

	this.Status = TASK_ONGOING
	session := CreateTaskSession(this.Id)
	this.sessions = append(this.sessions, session)
	this.currentSession = session

	return this.currentSession, nil
}

func (this *Task) Pause() (*TaskSession, error) {
	if !this.canPause() {
		return nil, fmt.Errorf("You can't pause this task")
	}

	this.Status = TASK_PAUSED
	session := this.getLastSession()
	session.End()
	this.currentSession = nil

	return session, nil
}

func (this *Task) End() (*TaskSession, error) {
	if !this.canEnd() {
		return nil, fmt.Errorf("You can't end this task")
	}

	this.Status = TASK_ENDED
	session := this.getLastSession()
	if session == nil {
		return nil, fmt.Errorf("[Impossible State] This task has no session!")
	}
	session.End()
	this.currentSession = nil

	return session, nil
}

func (this *Task) GetLastSessionDuration() time.Duration {
	lastSession := this.getLastSession()
	if lastSession != nil {
		return lastSession.Duration().Round(time.Second)
	}
	return time.Duration(0) * time.Second
}

func (this *Task) SetStatus(status string) {
	values := map[string]TaskStatus{
		"new":     TASK_NEW,
		"ongoing": TASK_ONGOING,
		"paused":  TASK_PAUSED,
		"ended":   TASK_ENDED,
	}

	st, ok := values[status]
	if ok {
		this.Status = st
	}

}

func (this *Task) SetSessions(sessions []*TaskSession) {
	this.sessions = sessions

}

func (this *Task) SetProject(id int64, name string) {
	this.Project = project.NewProject(id, name)
}

func (this *Task) GetProjectName() string {
	if this.Project == nil {
		return "N/A"
	}
	return this.Project.Name
}

func (this *Task) IsOngoing() bool {
	return this.Status == TASK_ONGOING
}

func (this *Task) getTotalElapsedTime() time.Duration {
	seconds := 0.0

	for _, s := range this.sessions {
		if s.isOngoing() {
			seconds += time.Now().Sub(s.StartsAt).Seconds()
		} else {
			seconds += s.Duration().Seconds()
		}
	}

	return time.Duration(seconds) * time.Second
}

func (this *Task) GetElapsedTime(mode TimerMode) time.Duration {
	sessionsLen := len(this.sessions)
	if sessionsLen == 0 {
		return time.Duration(0)
	}
	switch mode {
	case TIMER_MODE_SESSION:
		{
			session := this.sessions[sessionsLen-1]
			return time.Now().Sub(session.StartsAt)
		}
	case TIMER_MODE_TASK:
		{
			return this.getTotalElapsedTime()
		}
	}
	return time.Duration(0)
}

func (this *Task) GetSessions() []*TaskSession {
	return this.sessions
}
