package release

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/gogen"
)

func runPackr(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		fn := gogen.Install("github.com/gobuffalo/packr/v2/packr2", "-v")
		if err := fn(r); err != nil {
			return err
		}

		args := []string{"-v"}
		if opts.LegacyPackr {
			args = append(args, "--legacy")
		}
		return r.Exec(exec.Command("packr2", args...))
	}
}
