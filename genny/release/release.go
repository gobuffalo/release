package release

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	if _, err := exec.LookPath("git"); err != nil {
		return g, errors.New("git must be installed")
	}

	g.RunFn(WriteVersionFile(opts))

	g.RunFn(runShoulders)

	g.RunFn(runPackr)

	g.RunFn(makeInstall)

	g.RunFn(makeReleaseTest)

	g.RunFn(commit(opts))

	g.RunFn(tagRelease(opts))

	g.RunFn(runGoreleaser(opts))

	if len(opts.semVersion.Prerelease()) != 0 {
		g.RunFn(func(r *genny.Runner) error {
			r.Logger.Warn(preWarning)
			return nil
		})
	}
	return g, nil
}

const preWarning = `!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!

**THIS IS A PRE-RELEASE**

You MUST **MANUALLY** go to GitHub and edit the release accordingly!!!

NOTE: PRs welcome to make this happen automatically. :)
`

/*
* confirm GITHUB_TOKEN
* ask for version
* confirm semver version
* confirm branch
* (write version file)
* shoulders
* packr
* (make install)
* (make release-test)
* commit (if changes)
* tag
* push tags
* (goreleaser)
 */
