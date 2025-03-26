package cmd

import (
	"fmt"
	"strconv"

	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/spf13/cobra"
)

func newTaskUpdateCmd(ctx *Context) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update <task_id>",
		Args:  cobra.ExactArgs(1),
		Short: "Update a task",
		Long:  `Updates a the task name and assigned project.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			nFlag := cmd.Flag("name")
			newName := nFlag.Value.String()
			if len(newName) > 0 {
				err = updateTaskName(ctx, int64(taskId), nFlag.Value.String())
				if err != nil {
					return err
				}
			}

			pFlag := cmd.Flag("project")
			projectId, err := strconv.Atoi(pFlag.Value.String())
			if projectId > 0 {
				if err != nil {
					return err
				}
				err = assignToProject(ctx, int64(taskId), int64(projectId))
				if err != nil {
					return err
				}
			}

			fmt.Println("The task has been successfully updated! ✅")
			t, err := getTask(ctx, int64(taskId))
			ctx.TaskPrinter.PrintSingle(t)
			return nil
		},
	}
	updateCmd.Flags().Int64P("project", "p", 0, "The ID of the project you want this task to be part of")
	updateCmd.Flags().StringP("name", "n", "", "The new name of this task")
	updateCmd.MarkFlagsOneRequired("project", "name")
	return updateCmd
}

func getTask(ctx *Context, taskId int64) (*task.Task, error) {
	taskService := services.NewTaskService(ctx.DB)
	return taskService.GetTaskById(taskId, task.TaskFilter{})
}

func updateTaskName(ctx *Context, taskId int64, newName string) error {
	taskService := services.NewTaskService(ctx.DB)
	err := taskService.UpdateTaskName(taskId, newName)
	return err
}

func assignToProject(ctx *Context, taskId int64, projectId int64) error {

	projectService := services.NewProjectService(ctx.DB)
	project, err := projectService.GetProject(projectId)
	if err != nil {
		return err
	}
	taskService := services.NewTaskService(ctx.DB)
	err = taskService.AssignProject(taskId, project.Id)
	return err
}
