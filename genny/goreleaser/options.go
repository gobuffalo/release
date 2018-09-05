package goreleaser

import "errors"

type Options struct {
	Force    bool
	MainFile string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.MainFile) == 0 {
		return errors.New("goreleaser is only for binary applications")
	}
	return nil
}
