package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/release"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var releaseOptions = struct {
	*release.Options
	dryRun   bool
	yesToAll bool
}{
	Options: &release.Options{},
}

func init() {
	rootCmd.Flags().BoolVarP(&releaseOptions.dryRun, "dry-run", "d", false, "runs the release without actually releasing")
	rootCmd.Flags().BoolVarP(&releaseOptions.yesToAll, "yes", "y", false, "yes to all prompts")
	rootCmd.Flags().StringVarP(&releaseOptions.Version, "version", "v", "", "version you want to release")
	b, err := currentBranch()
	if err != nil {
		b = "master"
	}
	rootCmd.Flags().StringVarP(&releaseOptions.Branch, "branch", "b", b, "branch you want to use (default is current branch)")
	rootCmd.Flags().StringVarP(&releaseOptions.VersionFile, "version-file", "f", "", "write the version back into your version file")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "release",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := releaseOptions.Options
		if len(opts.Version) == 0 {
			var err error
			opts.Version, err = askForVersion()
			if err != nil {
				return errors.WithStack(err)
			}
		}

		var run *genny.Runner = genny.WetRunner(context.Background())
		if releaseOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		g, err := release.New(opts)
		if err != nil {
			return errors.WithStack(err)
		}
		run.With(g)

		err = confirm(fmt.Sprintf("are you sure you want to release %q (%s)?\n", opts.Version, opts.Branch), releaseOptions.yesToAll, run.Run)
		if err != nil {
			// attempt to remove the tag as something went wrong
			c := exec.Command("git", "push", "origin", ":"+opts.Version)
			if err := run.Exec(c); err != nil {
				run.Logger.Error(err)
			}

			c = exec.Command("git", "tag", "-d", opts.Version)
			if err := run.Exec(c); err != nil {
				run.Logger.Error(err)
			}
			return errors.WithStack(err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func askForVersion() (string, error) {
	cmd := exec.Command("git", "tag", "--sort", "-creatordate")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.WithStack(err)
	}

	lines := strings.Split(string(b), "\n")
	max := len(lines)
	if max > 5 {
		max = 5
	}
	for _, l := range lines[:max] {
		fmt.Println(l)
	}
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter version number (vx.x.x): ")
	v, err := r.ReadString('\n')
	if err != nil {
		return "", errors.WithStack(err)
	}
	v = strings.TrimSpace(v)
	return v, nil
}

func confirm(msg string, yes bool, fn func() error) error {
	if yes {
		return fn()
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(strings.ToLower(text))
	if text == "y" || text == "Y" {
		return fn()
	}
	return nil
}

func currentBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	b, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(b)), err
}
