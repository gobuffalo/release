package release

import (
	"context"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_writeVersionFile(t *testing.T) {

	table := []struct {
		vf  string
		pkg string
	}{
		{"foo/version.go", "foo"},
		{"version.go", "release"},
	}

	for _, tt := range table {
		t.Run(tt.vf, func(st *testing.T) {
			r := require.New(st)

			opts := &Options{
				VersionFile: tt.vf,
			}

			run := genny.DryRunner(context.Background())

			fn := writeVersionFile(opts)
			r.NoError(fn(run))

			res := run.Results()

			r.Len(res.Commands, 0)
			r.Len(res.Files, 1)

			f := res.Files[0]
			r.Equal(opts.VersionFile, f.Name())
			r.Contains(f.String(), "package "+tt.pkg)
		})
	}

}
