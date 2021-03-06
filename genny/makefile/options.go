package makefile

import (
	"fmt"
	"path/filepath"
)

type Options struct {
	Force       bool
	MainFile    string
	BuildPath   string
	VersionFile string
	Root        string
	Tags        []string
	WithPackr   bool
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
			return fmt.Errorf("%s is not a .go file", opts.MainFile)
		}
		opts.BuildPath = filepath.Dir(opts.MainFile)
		if len(opts.BuildPath) > 0 {
			opts.BuildPath = "./" + opts.BuildPath
		}
	}
	if len(opts.Root) == 0 {
		return fmt.Errorf("root can not be empty")
	}
	return nil
}
