package release

import (
	"bytes"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/shoulders/shoulders"
)

func runShoulders(r *genny.Runner) error {
	sh, err := shoulders.New()
	if err != nil {
		return err
	}
	in := &bytes.Buffer{}
	if err := sh.Write(in); err != nil {
		return err
	}
	f := genny.NewFile("SHOULDERS.md", in)

	return r.File(f)
}
