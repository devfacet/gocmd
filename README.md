# gocmd

[![Release][release-image]][release-url] [![Build Status][build-image]][build-url] [![Coverage][coverage-image]][coverage-url] [![Report][report-image]][report-url] [![GoDoc][doc-image]][doc-url]

A Go library for building command line applications.

## Features

- Advanced command line arguments handling
	- Subcommand handling
	- Short and long command line arguments
	- Multiple arguments (repeated or delimited)
	- Support for environment variables
	- Well formatted usage printing
	- Auto usage and version printing
	- Unknown argument handling
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
		Help      bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version   bool `short:"v" long:"version" description:"Display version"`
		VersionEx bool `long:"vv" description:"Display version (extended)"`
		Echo      struct {
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"echo" description:"Print arguments"`
		Math struct {
			Sqrt struct {
				Number float64 `short:"n" long:"number" required:"true" description:"Number"`
			} `command:"sqrt" description:"Calculate square root"`
			Pow struct {
				Base     float64 `short:"b" long:"base" required:"true" description:"Base"`
				Exponent float64 `short:"e" long:"exponent" required:"true" description:"Exponent"`
			} `command:"pow" description:"Calculate base exponential"`
		} `command:"math" description:"Math functions" nonempty:"true"`
	}{}

	// Echo command
	gocmd.HandleFlag("Echo", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Printf("%s\n", strings.Join(cmd.FlagArgs("Echo")[1:], " "))
		return nil
	})

	// Math commands
	gocmd.HandleFlag("Math.Sqrt", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println(math.Sqrt(flags.Math.Sqrt.Number))
		return nil
	})
	gocmd.HandleFlag("Math.Pow", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println(math.Pow(flags.Math.Pow.Base, flags.Math.Pow.Exponent))
		return nil
	})

	// Init the app
	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
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
  -h, --help         	Display usage
  -v, --version      	Display version
      --vv           	Display version (extended)

Commands:
  echo               	Print arguments
  math               	Math functions
    sqrt             	Calculate square root
      -n, --number   	Number
    pow              	Calculate base exponential
      -b, --base     	Base
      -e, --exponent 	Exponent

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
