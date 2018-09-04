<p align="center"><img src="https://github.com/gobuffalo/buffalo/blob/master/logo.svg" width="360"></p>

<p align="center">
<a href="https://godoc.org/github.com/gobuffalo/buffalo-release"><img src="https://godoc.org/github.com/gobuffalo/buffalo-release?status.svg" alt="GoDoc" /></a>
<a href="https://goreportcard.com/report/github.com/gobuffalo/buffalo-release"><img src="https://goreportcard.com/badge/github.com/gobuffalo/buffalo-release" alt="Go Report Card" /></a>
</p>

# github.com/gobuffalo/release

This tool is/will be used to release new tools from the Buffalo eco-system.

## Installation

```bash
$ go get -u -v github.com/gobuffalo/release
```

## Usage

```bash
$ release --help
```

The basics of what this command does are the following:

* confirm semver version and branch
* (write version file)
* shoulders
* packr
* (make install)
* (make release-test)
* (commit)
* tag release
* (goreleaser)

The items inside of `()` are only run if needed for that project.

## Doctor

The Doctor is in! Run this inside of the project you want to release **before** releasing it! It will save you time later on by making sure your system is ready release this particular project.

```bash
$ release doctor
```

