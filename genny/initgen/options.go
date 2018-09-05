package initgen

type Options struct {
	VersionFile string
	Version     string
	MainFile    string
	Force       bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.Version) == 0 {
		opts.Version = "v0.0.1"
	}
	if len(opts.VersionFile) == 0 {
		opts.Version = "version.go"
	}
	return nil
}
