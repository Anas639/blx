package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/anas639/blx/internal/tui"
	"github.com/spf13/cobra"
)

func newTaskTimeCmd(ctx *Context) *cobra.Command {
	var timeCmd = &cobra.Command{
		Use:   "time <task_id>",
		Short: "Track the elapsed time of a task in real-time.",
		Long: `The time command continuously displays the elapsed time in real-time.
By default it tracks the elapsed time since you started the task.
use the --session flag to track only the last session.
If you don't provide a <task_id> then the last active task will be selected.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskService := services.NewTaskService(ctx.DB)

			tsk, err := getTaskOrLast(args, taskService)
			if err != nil {
				return err
			}

			trackingMode := task.TIMER_MODE_TASK
			trackSession, _ := strconv.ParseBool(cmd.Flag("session").Value.String())
			if trackSession {
				trackingMode = task.TIMER_MODE_SESSION
			}

			isMinimalUI, _ := strconv.ParseBool(cmd.Flag("minimal").Value.String())
			if !isMinimalUI {
				ctx.TaskPrinter.PrintSingle(tsk)
			}

			elapsedTime := tsk.GetElapsedTime(trackingMode).Seconds()

			isPassiveMode, _ := strconv.ParseBool(cmd.Flag("passive").Value.String())
			if !isPassiveMode {
				timeTracker := tui.NewTrackerFromElapsed(elapsedTime)
				ch := timeTracker.Start()
				<-ch
			} else {
				fmt.Printf("%s", time.Duration(elapsedTime)*time.Second)
			}

			return nil
		},
	}

	timeCmd.Flags().BoolP("session", "s", false, "Track elapsed time of the last session only")
	timeCmd.Flags().BoolP("minimal", "m", false, "Print the elapsed time only without additional task details.")
	timeCmd.Flags().BoolP("passive", "p", false, "Will print the elapsed time and quit without keeping the time tracker active")

	return timeCmd
}

func getTaskOrLast(args []string, taskService *services.TaskService) (*task.Task, error) {
	if len(args) == 0 {
		return taskService.GetLastActiveTask()
	}

	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}

	return taskService.GetTaskById(int64(taskId), task.TaskFilter{
		Statuses: []task.TaskStatus{task.TASK_ONGOING},
	})

}
