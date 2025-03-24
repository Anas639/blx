package cmd

import "github.com/spf13/cobra"

const version = "0.0.1"

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version number",
	}
}
