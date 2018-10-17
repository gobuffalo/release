package initgen

import (
	"path/filepath"

	"github.com/pkg/errors"
)

type Options struct {
	VersionFile string
	Version     string
	MainFile    string
	Force       bool
	Root        string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
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
	}
	if len(opts.Root) == 0 {
		return errors.New("root can not be empty")
	}
	return nil
}
