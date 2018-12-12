package makefile

import (
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/packr/v2/plog"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	plog.Logger = logger.New(logger.DebugLevel)
	r := require.New(t)

	g, err := New(&Options{
		Root: ".",
	})
	r.NoError(err)

	run := gentest.NewRunner()
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	cmds := []string{"go get github.com/alecthomas/gometalinter", "gometalinter --install"}
	r.NoError(gentest.CompareCommands(cmds, res.Commands))

	files := []string{".gometalinter.json", "Makefile"}
	r.NoError(gentest.CompareFiles(files, res.Files))
}
