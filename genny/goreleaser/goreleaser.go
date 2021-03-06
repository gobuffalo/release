package goreleaser

import (
	"html/template"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/gogen"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	box := packr.New("release:genny:goreleaser", "../goreleaser/templates")
	if err := genny.ForceBox(g, box, opts.Force); err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}
	g.Transformer(gogen.TemplateTransformer(data, template.FuncMap{}))

	g.Transformer(genny.Dot())
	return g, nil
}
