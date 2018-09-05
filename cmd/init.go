package cmd

import (
	"context"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/goreleaser"
	"github.com/gobuffalo/release/genny/makefile"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var initOptions = struct {
	dryRun      bool
	force       bool
	mainFile    string
	versionFile string
}{}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "setups a project to use release",
	RunE: func(cmd *cobra.Command, args []string) error {
		var run *genny.Runner = genny.WetRunner(context.Background())
		if initOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		g, err := makefile.New(&makefile.Options{
			Force:       initOptions.force,
			VersionFile: initOptions.versionFile,
		})
		if err != nil {
			return errors.WithStack(err)
		}
		run.With(g)

		if len(initOptions.mainFile) != 0 {
			g, err = goreleaser.New(&goreleaser.Options{
				Force:    initOptions.force,
				MainFile: initOptions.mainFile,
			})
			if err != nil {
				return errors.WithStack(err)
			}
			run.With(g)
		}
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initOptions.dryRun, "dry-run", "d", false, "runs the generator dry")
	initCmd.Flags().BoolVarP(&initOptions.force, "force", "f", false, "force files to overwrite existing ones")
	initCmd.Flags().StringVarP(&initOptions.mainFile, "main-file", "m", "", "adds a .goreleaser.yml file (only for binary applications)")
	initCmd.Flags().StringVarP(&initOptions.versionFile, "version-file", "v", "", "path to a version file to maintain")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
