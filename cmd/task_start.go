package cmd

import (
	"strconv"

	"github.com/anas639/blx/internal/event"
	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/spf13/cobra"
)

func newTaskStartCmd(ctx *Context) *cobra.Command {

	var startCmd = &cobra.Command{
		Use:   "start <task_id>",
		Args:  cobra.ExactArgs(1),
		Short: "Start a task.",
		Long: `Marks a task as "ongoing" to begin tracking time.
If the task was paused, it resumes from where it left off.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			task, err := startTaskById(ctx, int64(taskId))
			if err != nil {
				return err
			}
			ctx.TaskPrinter.PrintSingle(task)
			return nil
		},
	}
	return startCmd
}

func startTaskById(ctx *Context, taskId int64) (*task.Task, error) {
	taskService := services.NewTaskService(ctx.DB)
	task, err := taskService.StartTask(int64(taskId))
	if err != nil {
		return nil, err
	}
	ctx.Broadcaster.SendEvent(event.NewPayload(event.EVENT_START, task.Id))
	return task, nil
}
