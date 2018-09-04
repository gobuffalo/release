package release

import (
	"bytes"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/shoulders/shoulders"
	"github.com/pkg/errors"
)

func runShoulders(r *genny.Runner) error {
	sh, err := shoulders.New()
	if err != nil {
		return errors.WithStack(err)
	}
	in := &bytes.Buffer{}
	if err := sh.Write(in); err != nil {
		return errors.WithStack(err)
	}
	f := genny.NewFile("shoulders.md", in)

	return r.File(f)
}
