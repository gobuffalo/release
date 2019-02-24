package cmd

import (
	"context"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/azure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var azureOptions = struct {
	*azure.Options
	dryRun bool
}{
	Options: &azure.Options{},
}

var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "generates azure pipelines files for Azure DevOps pipelines",
	RunE: func(cmd *cobra.Command, args []string) error {
		var run *genny.Runner = genny.WetRunner(context.Background())
		if azureOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		opts := azureOptions.Options
		if err := run.WithNew(azure.New(opts)); err != nil {
			return errors.WithStack(err)
		}

		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(azureCmd)
	azureCmd.Flags().BoolVarP(&azureOptions.dryRun, "dry-run", "d", false, "runs the generator dry")
	azureCmd.Flags().BoolVarP(&azureOptions.Force, "force", "f", false, "force files to overwrite existing ones")
}
