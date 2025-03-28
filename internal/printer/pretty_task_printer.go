package printer

import (
	"fmt"
	"os"

	"github.com/anas639/blx/internal/task"
	"github.com/jedib0t/go-pretty/v6/table"
)

type PrettyTaskPrinter struct {
}

func NewPrettyTaskPrinter() EntityPrinter[task.Task] {
	return &PrettyTaskPrinter{}
}

func (this *PrettyTaskPrinter) PrintSingle(tsk *task.Task) {
	this.PrintMany([]*task.Task{tsk})
}

func (this *PrettyTaskPrinter) PrintMany(tasks []*task.Task) {
	if len(tasks) == 0 {
		fmt.Println("You have no tasks 🍃")
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Status", "Duration", "Project"})
	for _, tsk := range tasks {
		t.AppendRow(table.Row{
			tsk.Id, tsk.Name, tsk.Status, tsk.GetElapsedTime(task.TIMER_MODE_TASK), tsk.GetProjectName(),
		})

	}
	t.Render()
}
