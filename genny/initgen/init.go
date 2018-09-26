package initgen

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/release/genny/git"
	"github.com/gobuffalo/release/genny/gomods"
	"github.com/gobuffalo/release/genny/goreleaser"
	"github.com/gobuffalo/release/genny/makefile"
	"github.com/gobuffalo/release/genny/release"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	g := genny.New()
	g.Box(packr.NewBox("../initgen/templates"))
	g.Transformer(genny.Dot())
	gg.Add(g)

	// set up git
	g, err := git.New(&git.Options{})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)

	// set up go mods if enabled
	g, err = gomods.New(&gomods.Options{})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)

	// write the version.go file
	g.RunFn(release.WriteVersionFile(&release.Options{
		VersionFile: opts.VersionFile,
		Version:     opts.Version,
	}))

	// write a new makefile
	g, err = makefile.New(&makefile.Options{
		Force:       opts.Force,
		VersionFile: opts.VersionFile,
		MainFile:    opts.MainFile,
	})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)

	// if there's a main file setup goreleaser
	if len(opts.MainFile) != 0 {
		g, err = goreleaser.New(&goreleaser.Options{
			Force:    opts.Force,
			MainFile: opts.MainFile,
		})
		if err != nil {
			return gg, errors.WithStack(err)
		}
		gg.Add(g)
	}

	// run go mod tidy again at the end
	g, err = gomods.New(&gomods.Options{})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)
	return gg, nil
}
