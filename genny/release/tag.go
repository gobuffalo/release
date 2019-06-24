package release

import (
	"os/exec"

	"github.com/gobuffalo/genny"
)

func tagRelease(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		c := exec.Command("git", "tag", opts.Version)
		if err := r.Exec(c); err != nil {
			return err
		}

		c = exec.Command("git", "push", "origin", opts.Branch)
		if err := r.Exec(c); err != nil {
			return err
		}

		c = exec.Command("git", "push", "origin", "--tags")
		if err := r.Exec(c); err != nil {
			return err
		}
		return nil
	}
}
