package cmd

import (
	"github.com/anas639/blx/internal/services"
	taskevent "github.com/anas639/blx/internal/task_event"
	"github.com/spf13/cobra"
)

func newWatchCmd(ctx *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Subscribe to active task events",
		Long: `Subscribes to active task events.
An active task is the last task that has been started.
It will print the elapsed time whenever a task is started.
If you pause a task or end it the timer will stop and a status message will be printed instead.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskService := services.NewTaskService(ctx.DB)
			ch, err := ctx.Listener.Listen()
			if err != nil {
				return err
			}
			consumer := taskevent.NewTaskEventConsumer(taskService, ch)
			consumer.Start()
			consumer.Wait()
			return nil
		},
	}
}
