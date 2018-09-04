package release

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func commit(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		check := exec.Command("git", "status")
		bb := &bytes.Buffer{}
		check.Stdout = bb
		check.Stderr = bb
		if err := r.Exec(check); err != nil {
			return errors.WithStack(err)
		}
		x := strings.TrimSpace(bb.String())
		if strings.Contains(x, "nothing to commit") {
			return nil
		}

		c := exec.Command("git", "add", ".")
		if err := r.Exec(c); err != nil {
			return errors.WithStack(err)
		}

		c = exec.Command("git", "commit", "-m", "version bump: "+opts.Version)
		if err := r.Exec(c); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}
