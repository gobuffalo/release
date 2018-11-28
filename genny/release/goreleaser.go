package release

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/git"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/pkg/errors"
)

var errFileNotFound = errors.New("file not found")

func releaserFile(r *genny.Runner) (genny.File, error) {
	if f, err := r.FindFile(".goreleaser.yml.plush"); err == nil {
		return f, nil
	}
	if f, err := r.FindFile(".goreleaser.yml"); err == nil {
		return f, nil
	}
	return nil, errFileNotFound
}

func runGoreleaser(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		gy, err := releaserFile(r)
		if err != nil {
			if errors.Cause(err) == errFileNotFound {
				r.Logger.Info("No .goreleaser.yml(.plush) detected so skipping goreleaser step")
				return nil
			}
			return errors.WithStack(err)
		}

		ctx := plush.NewContext()
		brew := true
		for _, x := range []string{"beta", "rc", "alpha"} {
			if strings.Contains(opts.Version, x) {
				brew = false
				break
			}
		}
		ctx.Set("brew", brew)
		t := plushgen.Transformer(ctx)
		f, err := t.Transform(gy)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := r.File(genny.NewFile(f.Name(), strings.NewReader(warningLabel+f.String()))); err != nil {
			return errors.WithStack(err)
		}

		if err := git.Run("add", ".goreleaser.yml")(r); err != nil {
			if errors.Cause(err) != git.ErrWorkingTreeClean {
				return errors.WithStack(err)
			}
		}
		if err := git.Run("commit", "-m", "generated goreleaser", ".goreleaser.yml")(r); err != nil {
			if errors.Cause(err) != git.ErrWorkingTreeClean {
				return errors.WithStack(err)
			}
		}

		if err := tagRelease(opts)(r); err != nil {
			return errors.WithStack(err)
		}

		c := exec.Command("goreleaser")
		if _, err := os.Stat(filepath.Join(".", "dist")); err == nil {
			c.Args = append(c.Args, "--rm-dist")
		}
		if err := r.Exec(c); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}

const warningLabel = "# Code generated by github.com/gobuffalo/release. DO NOT EDIT.\n# Edit .goreleaser.yml.plush instead\n\n"
