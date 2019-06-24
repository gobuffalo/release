package goreleaser

import (
	"fmt"
	"os/user"
	"path/filepath"
)

type Options struct {
	Force     bool
	MainFile  string
	BrewOwner string
	BrewTap   string
	Root      string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.MainFile) == 0 {
		return fmt.Errorf("goreleaser is only for binary applications")
	}
	if len(opts.BrewTap) == 0 {
		opts.BrewTap = "homebrew-tap"
	}

	if len(opts.Root) == 0 {
		return fmt.Errorf("root can not be empty")
	}

	if len(opts.BrewOwner) == 0 {
		user, err := user.Current()
		if err != nil {
			return err
		}
		opts.BrewOwner = user.Username
	}
	if len(opts.MainFile) > 0 {
		if filepath.Ext(opts.MainFile) != ".go" {
			return fmt.Errorf("%s is not a .go file", opts.MainFile)
		}
	}
	return nil
}
