package makefile

import (
	"strings"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{
		Root: ".",
	})
	r.NoError(err)

	run := gentest.NewRunner()
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 2)

	c := res.Commands[0]
	r.Equal("go get github.com/alecthomas/gometalinter", strings.Join(c.Args, " "))

	c = res.Commands[1]
	r.Equal("gometalinter --install", strings.Join(c.Args, " "))

	r.Len(res.Files, 2)

	f := res.Files[0]
	r.Equal(".gometalinter.json", f.Name())

	f = res.Files[1]
	r.Equal("Makefile", f.Name())
}
