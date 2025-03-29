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
			// by default we list ongoing tasks only
			statuses := []task.TaskStatus{task.TASK_ONGOING}

			shouldListAllStatuses, err := strconv.ParseBool(cmd.Flag("all").Value.String())

			if shouldListAllStatuses {
				statuses = task.AllStatuses()
			} else {
				statusFlag, err := cmd.Flags().GetStringSlice("status")
				if err == nil && len(statusFlag) > 0 {
					statuses = task.StatusesFromSlice(statusFlag)
				}
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

	lsCmd.Flags().BoolP("all", "a", false, "Lists all the tasks no matter the status")
	lsCmd.Flags().StringSliceP("status", "s", nil, "List tasks with status [new,ongoing,paused,ended]")
	return lsCmd
}
