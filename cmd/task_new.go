package cmd

import (
	"fmt"
	"strconv"

	"github.com/anas639/blx/internal/project"
	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

			// check project id
			projectFlag := cmd.Flag("project")
			if projectFlag.Value.String() != "" {
				project, err := assignTaskToProjectId(projectFlag, task.Id, taskService, services.NewProjectService(ctx.DB))
				if err != nil {
					fmt.Println(err.Error())
				} else {
					task.SetProject(project.Id, project.Name)
				}
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

func assignTaskToProjectId(flag *pflag.Flag, taskId int64, taskService *services.TaskService, projectService *services.ProjectService) (*project.Project, error) {
	projectId, err := strconv.Atoi(flag.Value.String())
	if err != nil {
		return nil, fmt.Errorf("Invalid project Id %s \n", flag.Value.String())
	}
	project, err := projectService.GetProjectById(int64(projectId))
	if err != nil {
		return nil, err
	}
	err = taskService.AssignProject(taskId, project.Id)
	if err != nil {
		return nil, err
	}
	return project, nil
}
