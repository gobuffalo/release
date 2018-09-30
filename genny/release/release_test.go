package release_test

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/release"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)
	opts := &release.Options{
		Version:     "v1.0.0",
		GitHubToken: "MYTOKEN",
		VersionFile: filepath.Join("foo", "version.go"),
		Branch:      "master",
	}
	g, err := release.New(opts)
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

func Test_New_Goreleaser(t *testing.T) {
	r := require.New(t)

	run := genny.DryRunner(context.Background())

	opts := &release.Options{
		Version:     "v1.0.0",
		GitHubToken: "MYTOKEN",
		VersionFile: filepath.Join("foo", "version.go"),
		Branch:      "master",
	}
	g, err := release.New(opts)
	r.NoError(err)

	run.Disk.Add(genny.NewFile(opts.VersionFile, strings.NewReader(versionFileBefore)))
	run.Disk.Add(genny.NewFile(".goreleaser.yml.plush", strings.NewReader(goReleaserTmpl)))

	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 11)
	c := res.Commands[10]
	r.Equal("goreleaser", strings.Join(c.Args, " "))

	r.Len(res.Files, 4)

	f := res.Files[0]
	r.Equal(".goreleaser.yml", f.Name())
	r.Contains(f.String(), "brew:")
}

func Test_New_Goreleaser_Beta(t *testing.T) {
	table := []struct {
		version string
		brew    bool
	}{
		{"v1.0.0", true},
		{"v1.0.0-beta.1", false},
		{"v1.0.0-rc.1", false},
	}

	for _, tt := range table {
		t.Run(tt.version, func(st *testing.T) {
			r := require.New(st)
			run := genny.DryRunner(context.Background())

			opts := &release.Options{
				Version:     tt.version,
				GitHubToken: "MYTOKEN",
				VersionFile: filepath.Join("foo", "version.go"),
				Branch:      "master",
			}
			g, err := release.New(opts)
			r.NoError(err)

			run.Disk.Add(genny.NewFile(opts.VersionFile, strings.NewReader(versionFileBefore)))
			run.Disk.Add(genny.NewFile(".goreleaser.yml.plush", strings.NewReader(goReleaserTmpl)))

			run.With(g)

			r.NoError(run.Run())

			res := run.Results()

			r.Len(res.Commands, 11)
			c := res.Commands[10]
			r.Equal("goreleaser", strings.Join(c.Args, " "))

			r.Len(res.Files, 4)

			f := res.Files[0]
			r.Equal(".goreleaser.yml", f.Name())
			if tt.brew {
				r.Contains(f.String(), "brew:")
			} else {
				r.NotContains(f.String(), "brew:")
			}

		})
	}

}

const versionFileBefore = `package foo

const Version = "development"
`

const versionFileAfter = `package foo

const Version = "v1.0.0"
`

const goReleaserTmpl = `builds:
-
  goos:
    - darwin
    - linux
    - windows
  env:
    - CGO_ENABLED=0

checksum:
  name_template: 'checksums.txt'

snapshot:
name_template: "{{"{{"}} .Tag {{"}}"}}-next"

changelog:
  sort: asc
  filters:
    exclude:
    - '^docs:'
    - '^test:'

<%= if (brew) { %>
brew:
  github:
    owner: {{ .opts.BrewOwner }}
    name: {{ .opts.BrewTap }}
<% } %>`
