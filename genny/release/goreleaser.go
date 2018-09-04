package release

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gobuffalo/genny"
)

func runGoreleaser(r *genny.Runner) error {
	if _, err := r.FindFile(".goreleaser.yml"); err != nil {
		r.Logger.Info("No .goreleaser.yml detected so skipping goreleaser step")
		return nil
	}
	c := exec.Command("goreleaser")
	if _, err := os.Stat(filepath.Join(".", "dist")); err == nil {
		c.Args = append(c.Args, "--rm-dist")
	}
	return r.Exec(c)
}
