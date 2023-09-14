package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	GitCommit = ""
	Version   = "development"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nGit commit: %s\n", Version, GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
