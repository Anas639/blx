package cmd

import (
	"strconv"

	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/tui"
	"github.com/spf13/cobra"
)

func newTaskDeleteCmd(ctx *Context) *cobra.Command {

	var deleteCmd = &cobra.Command{
		Use:   "delete <task_id>",
		Args:  cobra.ExactArgs(1),
		Short: "Delete a task.",
		Long: `Permanently removes a task and all associated time-tracking sessions.
This action cannot be undone!`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			// make sure the user did not execute this command accidentally
			userComfirms := tui.AskForConfirmatino("Are you sure you want to delete this task? All the session data will be lost!")
			if !userComfirms {
				return nil
			}

			taskService := services.NewTaskService(ctx.DB)
			err = taskService.DeleteTask(int64(taskId))

			if err != nil {
				return err
			}

			return nil
		},
	}

	return deleteCmd
}
