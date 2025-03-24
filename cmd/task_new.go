package cmd

import (
	"fmt"

	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newNewCmd(ctx *Context) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "new <task_name>",
		Args:  cobra.ExactArgs(1),
		Short: "Create a new task",
		Long: `Create a new task with the specified name.
Each task starts in the new state and can be assigned to a project.
To start tracking time use the --start flag or run blx start <id> 
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskName := args[0]
			taskService := services.NewTaskService(ctx.DB)
			task, err := taskService.CreateTask(taskName)

			if err != nil {
				return err
			}

			fmt.Println("Task successfully created ✅")
			ctx.TaskPrinter.PrintSingle(task)
			return nil
		},
	}
	newCmd.Flags().Int64P("project", "p", 0, "The ID of the project you want this task to be part of")
	newCmd.Flags().BoolP("start", "s", false, "Marks this task as ongoing")
	return newCmd
}
