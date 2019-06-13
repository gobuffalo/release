package initgen

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/release/genny/makefile"
	"github.com/pkg/errors"
)

type Options struct {
	*makefile.Options
	VersionFile string
	Version     string
	Force       bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.Root) == 0 {
		return errors.New("root can not be empty")
	}
	if len(opts.Version) == 0 {
		opts.Version = "v0.0.1"
	}
	if len(opts.VersionFile) == 0 {
		opts.Version = "version.go"
	}
	if len(opts.MainFile) > 0 {
		if filepath.Ext(opts.MainFile) != ".go" {
			return errors.Errorf("%s is not a .go file", opts.MainFile)
		}
	} else {
		if _, err := os.Stat(filepath.Join(opts.Root, "main.go")); err == nil {
			opts.MainFile = "main.go"
		}
	}
	return nil
}
