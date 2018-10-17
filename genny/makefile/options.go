package makefile

import (
	"path/filepath"

	"github.com/pkg/errors"
)

type Options struct {
	Force       bool
	MainFile    string
	BuildPath   string
	VersionFile string
	Root        string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.VersionFile) == 0 {
		opts.VersionFile = "version.go"
	}
	if len(opts.BuildPath) == 0 {
		opts.BuildPath = "."
	}
	if len(opts.MainFile) > 0 {
		if filepath.Ext(opts.MainFile) != ".go" {
			return errors.Errorf("%s is not a .go file", opts.MainFile)
		}
		opts.BuildPath = filepath.Dir(opts.MainFile)
		if len(opts.BuildPath) > 0 {
			opts.BuildPath = "./" + opts.BuildPath
		}
	}
	if len(opts.Root) == 0 {
		return errors.New("root can not be empty")
	}
	return nil
}
