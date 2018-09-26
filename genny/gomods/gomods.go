package gomods

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools/gomods"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	if !gomods.On() {
		return genny.New(), nil
	}

	g := genny.New()
	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}
	g.RunFn(func(r *genny.Runner) error {
		if _, err := r.FindFile("go.mod"); err == nil {
			return nil
		}
		return r.Exec(exec.Command("go", "mod", "init"))
	})
	g.Command(exec.Command("go", "mod", "tidy"))

	return g, nil
}
