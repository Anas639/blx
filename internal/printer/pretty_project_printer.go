package printer

import (
	"fmt"
	"os"

	"github.com/anas639/blx/internal/project"
	"github.com/jedib0t/go-pretty/v6/table"
)

type PrettyProjectPrinter struct {
}

func NewPrettyProjectPrinter() EntityPrinter[project.Project] {
	return &PrettyProjectPrinter{}
}

func (this *PrettyProjectPrinter) PrintSingle(prj *project.Project) {
	this.PrintMany([]*project.Project{prj})
}

func (this *PrettyProjectPrinter) PrintMany(projects []*project.Project) {
	if len(projects) == 0 {
		fmt.Println("You have no projects 🍃")
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name"})
	for _, project := range projects {
		t.AppendRow(table.Row{
			project.Id,
			project.Name,
		})

	}
	t.Render()
}
