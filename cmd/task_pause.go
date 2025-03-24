package cmd

import (
	"fmt"
	"strconv"

	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newTaskPauseCmd(ctx *Context) *cobra.Command {
	var pauseCmd = &cobra.Command{
		Use:   "pause <task_id>",
		Args:  cobra.ExactArgs(1),
		Short: "Pause a task.",
		Long: `Pauses an ongoing task and records a session entry.
Use "start <task_id>" to resume.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			taskService := services.NewTaskService(ctx.DB)
			task, err := taskService.PauseTask(int64(taskId))
			if err != nil {
				return err
			}

			fmt.Printf("[ ⏰ %s ] Task Paused ⏸️ \n", task.GetLastSessionDuration())
			ctx.TaskPrinter.PrintSingle(task)
			return nil
		},
	}
	return pauseCmd
}
