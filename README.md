# gocmd

[![Release][release-image]][release-url] [![Build Status][build-image]][build-url] [![Coverage][coverage-image]][coverage-url] [![Report][report-image]][report-url] [![GoDoc][doc-image]][doc-url]

A Go library for building command line applications.

## Features

- Advanced command line arguments handling
	- Short and long command line arguments
	- Subcommand handling
	- Well formatted usage printing
- Output tables in the terminal
- Template support for config files
- No external dependency

## Installation

```bash
go get github.com/devfacet/gocmd
```

## Usage

### A basic app

See [basic](examples/basic/main.go) for full code.

```go
func main() {
	flags := struct {
		Help      bool `short:"h" long:"help" default:"false" description:"Display usage"`
		Version   bool `short:"v" long:"version" default:"false" description:"Display version"`
		VersionEx bool `long:"vv" default:"false" description:"Display version (extended)"`
		Echo      struct {
		} `command:"echo" description:"Print arguments"`
	}{}

	app, err := gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags:       &flags,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Version
	if flags.Version || flags.VersionEx {
		app.PrintVersion(flags.VersionEx)
		return
	}

	// Echo
	if args, ok := app.LookupFlag("Echo"); ok && args != nil {
		fmt.Printf("%s\n", strings.TrimRight(strings.TrimLeft(fmt.Sprintf("%v", args[1:]), "["), "]"))
		return
	}

	// Help
	app.PrintUsage()
	return
}
```
```bash
cd examples/basic/
go build .
```
```
$ ./basic
Usage: basic [options...] COMMAND [options...]

A basic app

Options:       	
  -h, --help   	Display usage             	
  -v, --version	Display version           	
      --vv     	Display version (extended)	
               	
Commands:      	
  echo         	Print arguments 

```

## Build

```bash
go build .
```

## Test

```bash
./test.sh
```

## Release

```bash
git add CHANGELOG.md # update CHANGELOG.md
./release.sh v1.0.0  # replace "v1.0.0" with new version

git ls-remote --tags # check the new tag
```

## Contributing

- Code contributions must be through pull requests
- Run tests, linting and formatting before a pull request
- Pull requests can not be merged without being reviewed
- Use "Issues" for bug reports, feature requests and discussions
- Do not refactor existing code without a discussion
- Do not add a new third party dependency without a discussion
- Use semantic versioning and git tags for versioning

## License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.


[release-url]: https://github.com/devfacet/gocmd/releases/latest
[release-image]: https://img.shields.io/github/release/devfacet/gocmd.svg

[build-url]: https://travis-ci.org/devfacet/gocmd
[build-image]: https://travis-ci.org/devfacet/gocmd.svg?branch=master

[coverage-url]: https://coveralls.io/github/devfacet/gocmd?branch=master
[coverage-image]: https://coveralls.io/repos/devfacet/gocmd/badge.svg?branch=master&service=github

[report-url]: https://goreportcard.com/report/github.com/devfacet/gocmd
[report-image]: https://goreportcard.com/badge/github.com/devfacet/gocmd

[doc-url]: https://godoc.org/github.com/devfacet/gocmd
[doc-image]: https://godoc.org/github.com/devfacet/gocmd?status.svg
