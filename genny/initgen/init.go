package initgen

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/gogen/gomods"
	"github.com/gobuffalo/licenser/genny/licenser"
	"github.com/gobuffalo/release/genny/azure"
	"github.com/gobuffalo/release/genny/git"
	"github.com/gobuffalo/release/genny/goreleaser"
	"github.com/gobuffalo/release/genny/makefile"
	"github.com/gobuffalo/release/genny/release"
)

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, err
	}
	g := genny.New()
	g.Transformer(genny.Dot())
	gg.Add(g)

	// set up git
	g, err := git.New(&git.Options{
		Root: opts.Root,
	})
	if err != nil {
		return gg, err
	}
	gg.Add(g)

	if opts.Force {
		g = genny.New()
		g.RunFn(func(r *genny.Runner) error {
			for _, x := range []string{"go.mod", "go.sum"} {
				r.Delete(x)
			}
			return nil
		})
		gg.Add(g)
	}

	// set up go mods if enabled
	g, err = gomods.Init("", opts.Root)
	if err != nil {
		return gg, err
	}
	gg.Add(g)

	g, err = azure.New(&azure.Options{
		Force: opts.Force,
	})
	if err != nil {
		return gg, err
	}
	gg.Add(g)

	// write the version.go file
	g.RunFn(release.WriteVersionFile(&release.Options{
		VersionFile: opts.VersionFile,
		Version:     opts.Version,
	}))

	// write a new makefile
	g, err = makefile.New(opts.Options)
	if err != nil {
		return gg, err
	}
	gg.Add(g)

	// generate a license
	g, err = licenser.New(&licenser.Options{})
	if err != nil {
		return gg, err
	}
	gg.Add(g)

	// if there's a main file setup goreleaser
	if len(opts.MainFile) != 0 {
		g, err = goreleaser.New(&goreleaser.Options{
			Force:    opts.Force,
			MainFile: opts.MainFile,
			Root:     opts.Root,
		})
		if err != nil {
			return gg, err
		}
		gg.Add(g)
	}

	// run go mod tidy again at the end
	g, err = gomods.Tidy(opts.Root, false)
	if err != nil {
		return gg, err
	}
	gg.Add(g)
	return gg, nil
}
