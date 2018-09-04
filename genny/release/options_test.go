package release

import (
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/require"
)

func Test_Options_Validate(t *testing.T) {

	envy.Temp(func() {
		envy.Set("GITHUB_TOKEN", "")
		table := []struct {
			Options *Options
			Pass    bool
		}{
			{&Options{}, false},
			{&Options{GitHubToken: "foo", Version: "v1.0.0"}, true},
			{&Options{GitHubToken: "", Version: "v1.0.0"}, false},
			{&Options{GitHubToken: "foo", Version: ""}, false},
		}

		for _, tt := range table {
			t.Run(tt.Options.GitHubToken+"|"+tt.Options.Version, func(st *testing.T) {
				tt.Options.Branch = "master"
				r := require.New(st)
				if tt.Pass {
					r.NoError(tt.Options.Validate())
				} else {
					r.Error(tt.Options.Validate())
				}
			})
		}
	})
}
