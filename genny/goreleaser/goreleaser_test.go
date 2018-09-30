package goreleaser

import (
	"context"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{
		MainFile: "main.go",
	})
	r.NoError(err)

	run := genny.DryRunner(context.Background())
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 0)
	r.Len(res.Files, 1)

	f := res.Files[0]
	r.Equal(".goreleaser.yml.plush", f.Name())
}
