package git

type Options struct {
	Root string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
