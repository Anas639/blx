package cmd

import (
	"fmt"

	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newProjectNewCmd(ctx *Context) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "new <project_name>",
		Args:  cobra.ExactArgs(1),
		Short: "Create a new project.",
		Long: `Creates a project with the specific name.
Tasks can be assigned to projects for better oganization`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			projectService := services.NewProjectService(ctx.DB)
			project, err := projectService.CreateProject(projectName)

			if err != nil {
				return err
			}

			fmt.Printf("The project \"%s\" was successfully created \n", project.Name)
			ctx.ProjectPrinter.PrintSingle(project)
			return nil
		},
	}
	return newCmd
}
