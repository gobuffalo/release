package release

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/genny"
)

func makeInstall(r *genny.Runner) error {
	return makeRun("install", r)
}

func makeReleaseTest(r *genny.Runner) error {
	return makeRun("release-test", r)
}

func makeRun(target string, r *genny.Runner) error {
	if _, err := os.Stat(filepath.Join(r.Root, "Makefile")); err != nil {
		// No Makefile so we skip these steps
		r.Logger.Infof("No Makefile detected so skipping: make %s", target)
		return nil
	}
	bb := &bytes.Buffer{}
	out := io.MultiWriter(bb, os.Stdout)
	oerr := io.MultiWriter(bb, os.Stderr)
	cmd := exec.Command("make", target)
	cmd.Stdout = out
	cmd.Stderr = oerr

	if err := r.Exec(cmd); err != nil {
		if strings.Contains(bb.String(), fmt.Sprintf(noTarget, target)) {
			r.Logger.Infof("No target detected so skipping: make %s", target)
			return nil
		}
		return err
	}
	return nil
}

const noTarget = "No rule to make target `%s'"
