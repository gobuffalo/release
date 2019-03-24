package makefile

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	g.Root = opts.Root

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	box := packr.New("github.com/gobuffalo/release/genny/makefile/templates", "../makefile/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, errors.WithStack(err)
	}

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())

	return g, nil
}
