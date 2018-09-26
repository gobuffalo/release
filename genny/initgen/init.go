package initgen

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/goreleaser"
	"github.com/gobuffalo/release/genny/makefile"
	"github.com/gobuffalo/release/genny/release"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	g := genny.New()

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	g.RunFn(release.WriteVersionFile(&release.Options{
		VersionFile: opts.VersionFile,
		Version:     opts.Version,
	}))
	gg.Add(g)

	g, err := makefile.New(&makefile.Options{
		Force:       opts.Force,
		VersionFile: opts.VersionFile,
		MainFile:    opts.MainFile,
	})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)

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
	return gg, nil
}
