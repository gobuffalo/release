package release

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/gobuffalo/envy"
)

type Options struct {
	GitHubToken string
	Version     string
	Branch      string
	VersionFile string
	LegacyPackr bool
	SkipPackr   bool
	semVersion  *semver.Version
	// add your stuff here
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.GitHubToken) == 0 {
		opts.GitHubToken = envy.Get("GITHUB_TOKEN", "")
		if len(opts.GitHubToken) == 0 {
			return fmt.Errorf("you must set a GITHUB_TOKEN")
		}
	}
	if len(opts.Version) == 0 {
		opts.Version = "v0.0.1"
	}

	if !strings.HasPrefix(opts.Version, "v") {
		opts.Version = "v" + opts.Version
	}

	v, err := semver.NewVersion(strings.TrimPrefix(opts.Version, "v"))
	if err != nil {
		return err
	}
	opts.semVersion = v
	if len(opts.Branch) == 0 {
		opts.Branch = "master"
	}
	return nil
}
