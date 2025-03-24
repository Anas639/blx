package cmd

import (
	"strconv"

	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/spf13/cobra"
)

func newTaskLsCmd(ctx *Context) *cobra.Command {

	var lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all tasks",
		Long: `List all tasks with options to filter by status (new, ongoing, paused, ended) or by project.
Example: blx ls --status ongoing`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskService := services.NewTaskService(ctx.DB)
			shouldListAllStatuses, err := strconv.ParseBool(cmd.Flag("all").Value.String())
			statuses := []task.TaskStatus{}
			if !shouldListAllStatuses {
				statuses = append(statuses, task.TASK_ONGOING, task.TASK_PAUSED)
			}
			tasks, err := taskService.GetTasks(
				task.TaskFilter{
					Statuses: statuses,
				},
			)
			if err != nil {
				return err
			}
			ctx.TaskPrinter.PrintMany(tasks)
			return nil
		},
	}

	lsCmd.Flags().BoolP("all", "a", false, "Lists tasks with any status")
	return lsCmd
}
