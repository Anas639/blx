package cmd

import (
	"fmt"
	"strconv"

	"github.com/anas639/blx/internal/event"
	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/spf13/cobra"
)

func newTaskEndCmd(ctx *Context) *cobra.Command {

	var endCmd = &cobra.Command{
		Use:   "end <task_id>",
		Args:  cobra.ExactArgs(1),
		Short: "End a task.",
		Long: `Marks the task as "ended" and records a session entry.
Ended tasks cannot be resumed!`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			taskService := services.NewTaskService(ctx.DB)
			tsk, err := taskService.EndTask(int64(taskId))
			if err != nil {
				return err
			}
			fmt.Printf("[ ⏰ %s ] Task Successfully ended 🏁 \n", tsk.GetElapsedTime(task.TIMER_MODE_TASK))
			ctx.TaskPrinter.PrintSingle(tsk)
			ctx.Broadcaster.SendEvent(event.NewPayload(event.EVENT_END, tsk.Id))
			return nil
		},
	}
	return endCmd
}
