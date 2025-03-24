package cmd

import (
	"strconv"

	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/tui"
	"github.com/spf13/cobra"
)

func newProjectDeleteCmd(ctx *Context) *cobra.Command {

	var deleteCmd = &cobra.Command{
		Use:   "delete <project_id>",
		Args:  cobra.ExactArgs(1),
		Short: "Delete a project.",
		Long: `Permanently removes a project but it keeps all the associated tasks. 
This action cannot be undone!`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			// make sure the user did not execute this command accidentally
			userComfirms := tui.AskForConfirmatino("Are you sure you want to delete this project?")
			if !userComfirms {
				return nil
			}

			projectService := services.NewProjectService(ctx.DB)
			err = projectService.DeleteProject(int64(projectId))

			if err != nil {
				return err
			}

			return nil
		},
	}

	return deleteCmd
}
