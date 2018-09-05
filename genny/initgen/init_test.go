package initgen

import (
	"context"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	gg, err := New(&Options{
		VersionFile: "foo/bar/version.go",
		MainFile:    "./main.go",
	})
	r.NoError(err)

	run := genny.DryRunner(context.Background())
	run.WithGroup(gg)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 0)
	r.Len(res.Files, 3)

	f := res.Files[0]
	r.Equal(".goreleaser.yml", f.Name())

	f = res.Files[1]
	r.Equal("Makefile", f.Name())

	f = res.Files[2]
	r.Equal("foo/bar/version.go", f.Name())
	r.Contains(f.String(), `const Version = "v0.0.1"`)
}
