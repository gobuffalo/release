package release

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

var versionRx = regexp.MustCompile("[const|var] [vV]ersion = [`\"](.+)[`\"]")

func writeVersionFile(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {

		f, err := r.FindFile(opts.VersionFile)
		if err != nil {
			return errors.WithStack(err)
		}

		var matches []string
		bb := &bytes.Buffer{}
		for _, line := range strings.Split(f.String(), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "//") {
				bb.WriteString(line + "\n")
				continue
			}
			matches = versionRx.FindStringSubmatch(line)
			if len(matches) > 1 {
				v := matches[1]

				bb.WriteString(strings.Replace(line, v, opts.Version, 1) + "\n")
				continue
			}

			bb.WriteString(line + "\n")
		}

		body := strings.TrimSpace(bb.String())
		f = genny.NewFile(f.Name(), strings.NewReader(body+"\n"))
		return r.File(f)
	}
}
