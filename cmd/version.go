package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Current version of aws-cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("aws-cli/0.1.0")
		},
	}

	return cmd
}
