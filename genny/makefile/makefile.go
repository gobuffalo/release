package makefile

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}
	box := packr.NewBox("../makefile/templates")
	if err := g.Box(box); err != nil {
		return g, errors.WithStack(err)
	}

	err := box.Walk(func(path string, bf packr.File) error {
		f := genny.NewFile(path, bf)
		ff := genny.ForceFile(f, opts.Force)
		f, err := ff(f)
		if err != nil {
			return errors.WithStack(err)
		}
		g.File(f)
		return nil
	})
	if err != nil {
		return g, errors.WithStack(err)
	}
	return g, nil
}
