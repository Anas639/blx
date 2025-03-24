package services

import (
	"database/sql"
	"fmt"

	"github.com/anas639/blx/internal/project"
)

type ProjectService struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) *ProjectService {
	return &ProjectService{db: db}
}

func (this *ProjectService) CreateProject(name string) (*project.Project, error) {
	stmt, err := this.db.Prepare("INSERT INTO projects(name) VALUES(?)")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(name)
	if err != nil {
		return nil, err
	}
	projectID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	p := project.NewProject(projectID, name)
	return p, nil
}

func (this *ProjectService) GetProjects() ([]*project.Project, error) {
	projects := []*project.Project{}
	res, err := this.db.Query("SELECT id,name FROM projects")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var id int64
		var name string
		err = res.Scan(&id, &name)

		if err != nil {
			return nil, err
		}

		project := project.NewProject(id, name)
		projects = append(projects, project)
	}

	return projects, nil
}

func (this *ProjectService) GetProject(projectId int64) (*project.Project, error) {
	stmt, err := this.db.Prepare("select id,name from projects where id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(projectId)
	var id int64
	var name string
	err = row.Scan(&id, &name)
	if err != nil {
		return nil, fmt.Errorf("No project found with id %d", projectId)
	}
	p := project.NewProject(id, name)
	return p, nil
}

func (this *ProjectService) DeleteProject(projectId int64) error {
	stmt, err := this.db.Prepare("DELETE FROM projects WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(projectId)
	if err != nil {
		return err
	}
	if rcount, _ := res.RowsAffected(); rcount < 1 {
		return fmt.Errorf("No project found with id %d", projectId)
	}

	return nil
}

func (this *ProjectService) GetProjectById(projectId int64) (*project.Project, error) {
	stmt, err := this.db.Prepare("select p.id,p.name from projects p where p.id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(projectId)
	var id int64
	var name string
	err = row.Scan(&id, &name)
	if err != nil {
		return nil, fmt.Errorf("No project found with id %d", projectId)
	}
	p := project.NewProject(id, name)
	return p, nil
}

func (this *ProjectService) UpdateProjectName(projectId int64, newName string) error {
	stmt, err := this.db.Prepare("UPDATE projects SET name = ? where id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(newName, projectId)
	if rcount, _ := res.RowsAffected(); rcount < 1 {
		return fmt.Errorf("No project found with id %d", projectId)
	}
	return nil
}
