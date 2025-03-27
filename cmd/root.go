package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	command *cobra.Command
}

func (this *RootCmd) Execute(ctx context.Context) {
	err := this.command.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func NewRootCmd(ctx *Context) *RootCmd {
	var rootCmd = &cobra.Command{
		Use:     "blx",
		Version: version,
		Short:   "Track the time you spend on tasks and projects ⌛",
		Long: `Blx is a CLI tool that helps track the time you spend on tasks
while working on your computer 💻.
You can create, update, and delete tasks and projects, as well as manage tasks statuses.
Each time you pause or end a task a session is logged.
You can list tasks with filtering options and view project-related tasks.`,
	}

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	// task
	newCmd := newNewCmd(ctx)
	lsCmd := newTaskLsCmd(ctx)
	startCmd := newTaskStartCmd(ctx)
	pauseCmd := newTaskPauseCmd(ctx)
	endCmd := newTaskEndCmd(ctx)
	delCmd := newTaskDeleteCmd(ctx)
	updateCmd := newTaskUpdateCmd(ctx)
	timeCmd := newTaskTimeCmd(ctx)
	watchCmd := newWatchCmd(ctx)

	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(pauseCmd)
	rootCmd.AddCommand(endCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(watchCmd)

	// project
	prjCmd := newProjectCmd(ctx)
	rootCmd.AddCommand(prjCmd)

	return &RootCmd{
		command: rootCmd,
	}
}
