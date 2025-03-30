package cmd

import (
	"github.com/spf13/cobra"
)

func newProjectCmd(ctx *Context) *cobra.Command {
	var projectCmd = &cobra.Command{
		Use:   "project",
		Short: "Manage your projects",
	}

	createProject := newProjectNewCmd(ctx)
	lsProject := newProjectLsCmd(ctx)
	deleteProject := newProjectDeleteCmd(ctx)
	updateProject := newProjectUpdateCmd(ctx)
	findCmd := newProjectFindCmd(ctx)

	projectCmd.AddCommand(createProject)
	projectCmd.AddCommand(lsProject)
	projectCmd.AddCommand(deleteProject)
	projectCmd.AddCommand(updateProject)
	projectCmd.AddCommand(findCmd)

	return projectCmd
}
