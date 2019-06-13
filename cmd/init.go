package cmd

import (
	"context"
	"os"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/initgen"
	"github.com/gobuffalo/release/genny/makefile"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var initOptions = struct {
	*initgen.Options
	dryRun bool
}{
	Options: &initgen.Options{
		Options: &makefile.Options{},
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "setups a project to use release",
	RunE: func(cmd *cobra.Command, args []string) error {
		var run *genny.Runner = genny.WetRunner(context.Background())
		if initOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		opts := initOptions.Options
		pwd, _ := os.Getwd()
		opts.Root = pwd

		gg, err := initgen.New(opts)
		if err != nil {
			return errors.WithStack(err)
		}
		run.WithGroup(gg)

		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initOptions.dryRun, "dry-run", "d", false, "runs the generator dry")
	initCmd.Flags().BoolVarP(&initOptions.Force, "force", "f", false, "force files to overwrite existing ones")
	initCmd.Flags().BoolVarP(&initOptions.WithPackr, "with-packr", "p", false, "use packr2")
	initCmd.Flags().StringVarP(&initOptions.MainFile, "main-file", "m", "", "adds a .goreleaser.yml file (only for binary applications)")
	initCmd.Flags().StringVarP(&initOptions.VersionFile, "version-file", "v", "version.go", "path to a version file to maintain")
	initCmd.Flags().StringSliceVarP(&initOptions.Tags, "tags", "t", []string{}, "tags for the Makefile")
}
