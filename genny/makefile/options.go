package makefile

type Options struct {
	Force       bool
	VersionFile string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
