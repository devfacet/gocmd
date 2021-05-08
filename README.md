# gocmd

[![Release][release-image]][release-url] [![GoDoc][doc-image]][doc-url] [![Report][report-image]][report-url]

A Go library for building command line applications.

> Note that the terminal commands are tested on MacOS and should work on POSIX compatible systems.

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

```shell
go get github.com/devfacet/gocmd/v3
```

## Usage

### A basic app

For the full code [click here](examples/basic/main.go).

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
		Description: "A basic app",
		Version:     fmt.Sprintf("%s (%s)", version, gitCommit),
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
```
```shell
# Build the example app:
make build

# Run it:
gocmdbasic
```
```shell
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

## Test

```sh
# Test everything:
make test

# For BDD development:
# It will open a new browser window. Make sure:
#   1. There is no errors on the terminal window.
#   2. There is no other open GoConvey page.
make test-ui
```

## Release

```sh
git add CHANGELOG.md # update CHANGELOG.md first
make release v1.0.0  # replace "v1.0.0" with new version
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.


[doc-url]: https://pkg.go.dev/github.com/devfacet/gocmd/v3
[doc-image]: https://godoc.org/github.com/devfacet/gocmd?status.svg

[release-url]: https://github.com/devfacet/gocmd/releases/latest
[release-image]: https://img.shields.io/github/release/devfacet/gocmd.svg?style=flat-square

[report-url]: https://goreportcard.com/report/github.com/devfacet/gocmd
[report-image]: https://goreportcard.com/badge/github.com/devfacet/gocmd?style=flat-square
