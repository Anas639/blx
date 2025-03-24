package project

import "fmt"

type Project struct {
	Id   int64
	Name string
}

func CreateProject(name string) *Project {
	return &Project{Name: name}
}

func NewProject(id int64, name string) *Project {
	return &Project{
		Id:   id,
		Name: name,
	}
}

func (this *Project) String() string {
	return fmt.Sprintf("#%d %s", this.Id, this.Name)
}
