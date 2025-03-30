package cmd

import (
	"github.com/anas639/blx/internal/services"
	"github.com/spf13/cobra"
)

func newProjectFindCmd(ctx *Context) *cobra.Command {
	findCmd := &cobra.Command{
		Use:  "find",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			keyword := args[0]
			projectService := services.NewProjectService(ctx.DB)
			projects, err := projectService.Search(keyword)
			if err != nil {
				return err
			}
			ctx.ProjectPrinter.PrintMany(projects)
			return nil
		},
	}
	return findCmd
}
