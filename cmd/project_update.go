package cmd

import (
	"fmt"
	"strconv"

	"github.com/anas639/blx/internal/project"
	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newProjectUpdateCmd(ctx *Context) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update <project-id>",
		Args:  cobra.ExactArgs(1),
		Short: "Update a task",
		Long:  `Updates the project name.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectId, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			nFlag := cmd.Flag("name")
			newName := nFlag.Value.String()
			if len(newName) > 0 {
				err = updateProjectName(ctx, int64(projectId), nFlag.Value.String())
				if err != nil {
					return err
				}
			}

			fmt.Println("The project has been successfully updated! ✅")
			t, err := getProject(ctx, int64(projectId))
			ctx.ProjectPrinter.PrintSingle(t)
			return nil
		},
	}
	updateCmd.Flags().StringP("name", "n", "", "The new name of this project")
	updateCmd.MarkFlagsOneRequired("name")
	return updateCmd
}

func getProject(ctx *Context, projectId int64) (*project.Project, error) {
	projectService := services.NewProjectService(ctx.DB)
	return projectService.GetProjectById(projectId)
}

func updateProjectName(ctx *Context, projectId int64, newName string) error {
	projectService := services.NewProjectService(ctx.DB)
	err := projectService.UpdateProjectName(projectId, newName)
	return err
}
