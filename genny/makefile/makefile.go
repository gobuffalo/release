package makefile

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	g.Root = opts.Root

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	box := packr.New("release:genny:makefile", "../makefile/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, errors.WithStack(err)
	}

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())
	g.RunFn(gotools.Install("github.com/alecthomas/gometalinter"))

	g.RunFn(func(r *genny.Runner) error {
		c := exec.Command("gometalinter", "--install")
		return r.Exec(c)
	})

	return g, nil
}
