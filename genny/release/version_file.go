package release

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var versionRx = regexp.MustCompile("[const|var] [vV]ersion = ([`\"].*[`\"])")

func WriteVersionFile(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {

		if len(opts.Version) == 0 {
			return errors.New("version can not be blank")
		}

		f, err := r.FindFile(opts.VersionFile)
		if err != nil {
			f, err = defaultVersionFile(opts.VersionFile)
			if err != nil {
				return errors.WithStack(err)
			}
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

				bb.WriteString(strings.Replace(line, v, `"`+opts.Version+`"`, 1) + "\n")
				continue
			}

			bb.WriteString(line + "\n")
		}

		body := strings.TrimSpace(bb.String())
		f = genny.NewFile(f.Name(), strings.NewReader(body+"\n"))
		return r.File(f)
	}
}

func defaultVersionFile(name string) (genny.File, error) {
	dir := filepath.Dir(name)
	files, err := ioutil.ReadDir(dir)
	_, ok := err.(*os.PathError)
	if err != nil && !ok {
		return nil, errors.WithStack(err)
	}
	var pkg string
	if len(files) == 0 {
		pkg = filepath.Base(dir)
	} else {
		for _, fi := range files {
			ext := filepath.Ext(fi.Name())
			if ext != ".go" {
				continue
			}
			if strings.HasSuffix(fi.Name(), "_test.go") {
				continue
			}
			b, err := ioutil.ReadFile(fi.Name())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			xf := genny.NewFile(fi.Name(), bytes.NewReader(b))
			pkg, err = gotools.PackageName(xf)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			break
		}
	}

	f := genny.NewFile(name+".plush", strings.NewReader(versionTmpl))

	if len(pkg) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			return f, errors.WithStack(err)
		}
		pkg = filepath.Base(pwd)
	}

	ctx := plush.NewContext()
	ctx.Set("pkg", pkg)
	t := plushgen.Transformer(ctx)
	f, err = t.Transform(f)
	if err != nil {
		return f, errors.WithStack(err)
	}
	return f, nil
}

const versionTmpl = `package <%= pkg %>

const Version = ""`
