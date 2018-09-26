package gomods

import (
	"context"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools/gomods"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{})
	r.NoError(err)

	run := genny.DryRunner(context.Background())
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	if !gomods.On() {
		r.Len(res.Commands, 0)
		r.Len(res.Files, 0)
		return
	}
	r.Len(res.Commands, 2)
	r.Len(res.Files, 0)
}
