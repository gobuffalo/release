package goreleaser

type Options struct {
	Force bool
	Brew  bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
