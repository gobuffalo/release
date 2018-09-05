package makefile

type Options struct {
	Force       bool
	VersionFile string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.VersionFile) == 0 {
		opts.VersionFile = "version.go"
	}
	return nil
}
