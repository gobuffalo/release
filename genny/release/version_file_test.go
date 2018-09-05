package release

import (
	"context"
	"fmt"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_writeVersionFile(t *testing.T) {

	table := []struct {
		vf  string
		pkg string
		v   string
	}{
		{"foo/version.go", "foo", "v1.0.0"},
		{"version.go", "release", "v2.0.0"},
	}

	for _, tt := range table {
		t.Run(tt.vf, func(st *testing.T) {
			r := require.New(st)

			opts := &Options{
				VersionFile: tt.vf,
				Version:     tt.v,
			}

			run := genny.DryRunner(context.Background())

			fn := WriteVersionFile(opts)
			r.NoError(fn(run))

			res := run.Results()

			r.Len(res.Commands, 0)
			r.Len(res.Files, 1)

			f := res.Files[0]
			r.Equal(opts.VersionFile, f.Name())
			body := f.String()
			r.Contains(body, "package "+tt.pkg)
			r.Contains(body, fmt.Sprintf(`const Version = "%s"`, tt.v))
		})
	}

}
