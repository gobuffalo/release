package git

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	g.Root = opts.Root

	g.Command(exec.Command("git", "init"))
	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("release:genny/git", "../git/templates")); err != nil {
		return g, err
	}
	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())
	return g, nil
}
