package release

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/pkg/errors"
)

func runPackr(r *genny.Runner) error {
	fn := gotools.Install("github.com/gobuffalo/packr/packr", "-v")
	if err := fn(r); err != nil {
		return errors.WithStack(err)
	}
	return r.Exec(exec.Command("packr", "-v"))
}
