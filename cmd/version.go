package cmd

import (
	"fmt"

	"github.com/gobuffalo/release/release"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of release",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("release", release.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
