package goreleaser

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	box := packr.NewBox("../goreleaser/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, errors.WithStack(err)
	}
	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())
	return g, nil
}
