package makefile

type Options struct {
	Force bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
