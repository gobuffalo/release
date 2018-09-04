package release

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)
	opts := &Options{
		Version:     "v1.0.0",
		GitHubToken: "MYTOKEN",
		VersionFile: filepath.Join("foo", "version.go"),
		Branch:      "master",
	}
	g, err := New(opts)
	r.NoError(err)

	run := genny.DryRunner(context.Background())
	run.Disk.Add(genny.NewFile(opts.VersionFile, strings.NewReader(versionFileBefore)))
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 8)

	c := res.Commands[0]
	r.Equal("go get -v github.com/gobuffalo/packr/packr", strings.Join(c.Args, " "))

	c = res.Commands[1]
	r.Equal("packr -v", strings.Join(c.Args, " "))

	r.Len(res.Files, 2)

	f := res.Files[0]
	r.Equal("foo/version.go", f.Name())
	body, err := ioutil.ReadAll(f)
	r.NoError(err)
	r.Equal(strings.TrimSpace(versionFileAfter), strings.TrimSpace(string(body)))

	f = res.Files[1]
	r.Equal("shoulders.md", f.Name())
	r.Contains(f.String(), "Stands on the Shoulders of Giants")
}

const versionFileBefore = `package foo

const Version = "development"
`

const versionFileAfter = `package foo

const Version = "v1.0.0"
`
