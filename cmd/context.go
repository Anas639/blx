package cmd

import (
	"database/sql"

	"github.com/anas639/blx/internal/printer"
	"github.com/anas639/blx/internal/project"
	"github.com/anas639/blx/internal/task"
)

type Context struct {
	DB             *sql.DB
	TaskPrinter    printer.EntityPrinter[task.Task]
	ProjectPrinter printer.EntityPrinter[project.Project]
}
