package azure

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}
	box := packr.New("github.com/gobuffalo/release/genny/azure/templates", "../azure/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, err
	}
	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}
