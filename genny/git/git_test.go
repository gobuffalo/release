package git

import (
	"strings"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{})
	r.NoError(err)

	run := gentest.NewRunner()
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 1)

	c := res.Commands[0]
	r.Equal("git init", strings.Join(c.Args, " "))

	r.Len(res.Files, 1)

	f := res.Files[0]
	r.Equal(".gitignore", f.Name())
}
