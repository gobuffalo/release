<p align="center"><img src="https://github.com/gobuffalo/buffalo/blob/master/logo.svg" width="360"></p>

<p align="center">
<a href="https://godoc.org/github.com/gobuffalo/release"><img src="https://godoc.org/github.com/gobuffalo/release?status.svg" alt="GoDoc" /></a>
<a href="https://dev.azure.com/markbates/buffalo/_build?definitionId=3"><img src="https://dev.azure.com/markbates/buffalo/_apis/build/status/gobuffalo.release?branchName=master" alt="CI" /></a>
<a href="https://goreportcard.com/report/github.com/gobuffalo/release"><img src="https://goreportcard.com/badge/github.com/gobuffalo/release" alt="Go Report Card" /></a>
</p>

# github.com/gobuffalo/release

This tool is/will be used to release new tools from the Buffalo eco-system. (Although it should work for everyone.)

---

## Installation

```bash
$ go get -u -v github.com/gobuffalo/release
```

## Setting up Your Project

Release is designed to work with any Go project using GitHub.com as it's VCS (Pull Requests welcome), so there isn't any setup needed.

Should you want to make your "releasing" experience a little easier, then the `release init` command is for you.

```bash
$ release init -h

setups a project to use release

Usage:
  release init [flags]

Flags:
  -d, --dry-run               runs the generator dry
  -f, --force                 force files to overwrite existing ones
  -h, --help                  help for init
  -m, --main-file string      adds a .goreleaser.yml file (only for binary applications)
  -v, --version-file string   path to a version file to maintain (default "version.go")
```

### Flags

#### `-m` - (Binary Applications Only)

If you pass `-m path/to/main.go` then Release will generate a [`.goreleaser.yml`](https://goreleaser.com) file for you.

#### `-v` - Version File

Release will automatically manage a "version" file for you. If you have an existing version file the `-v` allows Release to automatically update that file for you when you run the `make release` task.

If don't pass a `-v` flag then Release will create a new file for you, `./version.go`.

```go
package mypackage

const Version = "v1.0.0"
```

---

## Releasing

The basics of what this command does are the following:

* Confirm semver version and branch
* (write version file)
* Generate shoulders.md file
* Run `packr2` to pack any boxes
* (make install)
* (make release-test)
* (commit version bump)
* tag release
* (run goreleaser)

The items inside of `()` are only run if needed for that project.

```bash
$ release -h

Usage:
  release [flags]
  release [command]

Available Commands:
  doctor      checks to make sure your system is ready to release
  help        Help about any command
  init        setups a project to use release
  version     current version of release

Flags:
  -b, --branch string         branch you want to use (default is current branch) (default "master")
  -d, --dry-run               runs the release without actually releasing
  -h, --help                  help for release
  -v, --version string        version you want to release
  -f, --version-file string   write the version back into your version file (default "version.go")
  -y, --yes                   yes to all prompts

Use "release [command] --help" for more information about a command.
```

### Flags

#### `-v` - Version Flag

If a `-v v1.0.0` flag is passed then the version from that flag will be used. If no `-v` flag is used, then a prompt will be presented to the user, displaying up to the last file release tags (for context), and the user can enter a tag at the prompt.

```bash
$ release

v2.0.11
v2.0.10
v2.0.9
v2.0.8
v2.0.7
Enter version number (vx.x.x):
```

#### `-f` - Version File

Release will automatically manage a "version" file for you. If you have an existing version file the `-f` allows Release to automatically update that file for you when releasing.

If don't pass a `-f` flag then Release will create a new file for you, `./version.go`.

```go
package mypackage

const Version = "v1.0.0"
```

---

## Doctor

The Doctor is in! Run this inside of the project you want to release **before** releasing it! It will save you time later on by making sure your system is ready release this particular project.

```bash
$ release doctor
```

