package cmd

import (
	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newProjectLsCmd(ctx *Context) *cobra.Command {

	var lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all project",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectServic := services.NewProjectService(ctx.DB)
			projects, err := projectServic.GetProjects()
			if err != nil {
				return err
			}
			ctx.ProjectPrinter.PrintMany(projects)
			return nil
		},
	}

	return lsCmd
}
