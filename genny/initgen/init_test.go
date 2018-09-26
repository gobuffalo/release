package initgen

import (
	"context"
	"strings"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools/gomods"
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

	var cmds []string
	res := run.Results()
	if !gomods.On() {
		cmds = []string{"git init", "go get github.com/alecthomas/gometalinter", "gometalinter --install"}
	} else {
		cmds = []string{"git init", "go mod init", "go mod tidy", "go get github.com/alecthomas/gometalinter", "gometalinter --install"}
	}

	r.Len(res.Commands, len(cmds))
	for i, x := range cmds {
		r.Equal(x, strings.Join(res.Commands[i].Args, " "))
	}

	r.Len(res.Files, 5)

	f := res.Files[0]
	r.Equal(".gitignore", f.Name())

	f = res.Files[1]
	r.Equal(".gometalinter.json", f.Name())

	f = res.Files[2]
	r.Equal(".goreleaser.yml", f.Name())

	f = res.Files[3]
	r.Equal("Makefile", f.Name())

	f = res.Files[4]
	r.Equal("foo/bar/version.go", f.Name())
	r.Contains(f.String(), `const Version = "v0.0.1"`)
}
