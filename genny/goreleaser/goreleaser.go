package goreleaser

import (
	"html/template"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/packr"
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

	data := map[string]interface{}{
		"opts": opts,
	}
	g.Transformer(gotools.TemplateTransformer(data, template.FuncMap{}))

	g.Transformer(genny.Dot())
	return g, nil
}
