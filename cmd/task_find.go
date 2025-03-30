package cmd

import (
	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/spf13/cobra"
)

func newTaskFindCmd(ctx *Context) *cobra.Command {
	findCmd := &cobra.Command{
		Use:  "find",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			keyword := args[0]
			taskServise := services.NewTaskService(ctx.DB)
			tasks, err := taskServise.Search(keyword, task.TaskFilter{})
			if err != nil {
				return err
			}
			ctx.TaskPrinter.PrintMany(tasks)
			return nil
		},
	}
	return findCmd
}
