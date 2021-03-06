package makefile

import (
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	g.Root = opts.Root

	if err := opts.Validate(); err != nil {
		return g, err
	}

	box := packr.New("github.com/gobuffalo/release/genny/makefile/templates", "../makefile/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, err
	}

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	ctx.Set("tags", strings.Join(opts.Tags, " "))
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())

	return g, nil
}
