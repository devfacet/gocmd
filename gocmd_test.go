/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocmd_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/devfacet/gocmd"
	. "github.com/smartystreets/goconvey/convey"
)

var osArgs []string

func init() {
	osArgs = make([]string, len(os.Args))
	copy(osArgs, os.Args)
}

func resetArgs() {
	os.Args = make([]string, len(osArgs))
	copy(os.Args, osArgs)
}

func TestNew(t *testing.T) {
	Convey("should create a new command", t, func() {
		resetArgs()
		os.Args = append(os.Args, "-f=true", "--bar=test", "-b", "baz", "test", "-f=false", "-b=bar")

		flags := struct {
			Foo     bool   `short:"f"`
			Bar     string `long:"bar"`
			Baz     string `short:"b" long:"baz"`
			Command struct {
				Foo bool   `short:"f"`
				Bar string `short:"b"`
			} `command:"test"`
		}{}
		cmd, err := gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags:       &flags,
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.Name(), ShouldEqual, "test")
		So(cmd.Version(), ShouldEqual, "1.0.0")
		So(cmd.Description(), ShouldEqual, "Test")
		So(flags.Foo, ShouldEqual, true)
		So(flags.Bar, ShouldEqual, "test")
		So(flags.Baz, ShouldEqual, "baz")
		So(flags.Command.Foo, ShouldEqual, false)
		So(flags.Command.Bar, ShouldEqual, "bar")

		resetArgs()
	})

	Convey("should fail to create a new command", t, func() {
		resetArgs()

		cmd, err := gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"foo"`
			}{},
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("short argument foo in Foo field must be one character long"))
		So(cmd, ShouldBeNil)

		cmd, err = gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"f" required:"true"`
			}{},
			AnyError: true,
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("argument -f is required"))
		So(cmd, ShouldBeNil)

		cmd, err = gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"f"`
				Bar bool `short:"f"`
			}{},
			Logger:      log.New(ioutil.Discard, "", 0),
			ExitOnError: true,
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("short argument f in Bar field is already defined in Foo field"))
		So(cmd, ShouldBeNil)

		cmd, err = gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"f" required:"true"`
			}{},
			Logger:      log.New(ioutil.Discard, "", 0),
			ExitOnError: true,
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("argument -f is required"))
		So(cmd, ShouldBeNil)
	})
}

func TestCmd_Name(t *testing.T) {
	Convey("should return the correct command name", t, func() {
		cmd, err := gocmd.New(gocmd.Options{
			Name: "test",
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.Name(), ShouldEqual, "test")
	})
}

func TestCmd_Version(t *testing.T) {
	Convey("should return the correct command version", t, func() {
		cmd, err := gocmd.New(gocmd.Options{
			Version: "1.0.0",
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.Version(), ShouldEqual, "1.0.0")
	})
}

func TestCmd_Description(t *testing.T) {
	Convey("should return the correct command description", t, func() {
		cmd, err := gocmd.New(gocmd.Options{
			Description: "Test",
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.Description(), ShouldEqual, "Test")
	})
}

func TestCmd_LookupFlag(t *testing.T) {
	Convey("should lookup a flag", t, func() {
		resetArgs()
		os.Args = append(os.Args, "-f=true")

		cmd, err := gocmd.New(gocmd.Options{
			Flags: &struct {
				Foo bool `short:"f" default:"true"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		v, ok := cmd.LookupFlag("Foo")
		So(v, ShouldContain, "true")
		So(ok, ShouldEqual, true)
		v, ok = cmd.LookupFlag("Bar")
		So(v, ShouldBeNil)
		So(ok, ShouldEqual, false)

		resetArgs()
	})
}

func TestCmd_FlagValue(t *testing.T) {
	Convey("should return the correct flag value", t, func() {
		resetArgs()
		os.Args = append(os.Args, "-f=true")

		cmd, err := gocmd.New(gocmd.Options{
			Flags: &struct {
				Foo bool `short:"f" default:"true"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		v, ok := cmd.FlagValue("Foo").(bool)
		So(v, ShouldEqual, true)
		So(ok, ShouldEqual, true)
		v, ok = cmd.FlagValue("Bar").(bool)
		So(v, ShouldEqual, false)
		So(ok, ShouldEqual, false)

		resetArgs()
	})
}

func TestCmd_FlagArgs(t *testing.T) {
	Convey("should return the correct flag args", t, func() {
		resetArgs()
		os.Args = append(os.Args, "-f=true")

		cmd, err := gocmd.New(gocmd.Options{
			Flags: &struct {
				Foo bool `short:"f" default:"true"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		args := cmd.FlagArgs("Foo")
		So(args, ShouldHaveLength, 1)
		args = cmd.FlagArgs("Bar")
		So(args, ShouldHaveLength, 0)

		resetArgs()
	})
}

func TestCmd_FlagErrors(t *testing.T) {
	Convey("should return the flag errors", t, func() {
		cmd, err := gocmd.New(gocmd.Options{
			Flags: &struct {
				Foo struct {
				} `command:"foo" global:"true"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		flagErrors := cmd.FlagErrors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("command foo can't be global"))
	})
}

func TestHandleFlag(t *testing.T) {
	Convey("should fail to add flag handler", t, func() {
		fh, err := gocmd.HandleFlag("", func(cmd *gocmd.Cmd, args []string) error {
			return nil
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("invalid flag name"))
		So(fh, ShouldBeNil)
	})

	Convey("should add flag handler", t, func() {
		resetArgs()

		os.Args = []string{"gocmd.test", "fh1", "fh2"}
		fhCnt := 0
		fh, _ := gocmd.HandleFlag("FH1", func(cmd *gocmd.Cmd, args []string) error {
			fhCnt++
			return errors.New("handler error")
		})
		fh.SetPriority(1)
		fh.SetExitOnError(true)
		gocmd.HandleFlag("FH2", func(cmd *gocmd.Cmd, args []string) error {
			fhCnt++
			return nil
		})
		fh.SetPriority(2)
		cmd, err := gocmd.New(gocmd.Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				FH1 struct{} `command:"fh1"`
				FH2 struct{} `command:"fh2"`
			}{},
			Logger: log.New(ioutil.Discard, "", 0),
		})
		So(err, ShouldNotBeNil)
		So(err, ShouldBeError, errors.New("handler error"))
		So(cmd, ShouldBeNil)
		So(fhCnt, ShouldEqual, 2)

		resetArgs()
	})
}

func ExampleNew_usage() {
	os.Args = []string{"gocmd.test"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
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
			} `command:"math" description:"Math functions"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// Usage: basic [options...] COMMAND [options...]
	//
	// A basic app
	//
	// Options:
	//   -h, --help         	Display usage
	//   -v, --version      	Display version
	//       --vv           	Display version (extended)
	//
	// Commands:
	//   echo               	Print arguments
	//   math               	Math functions
	//     sqrt             	Calculate square root
	//       -n, --number   	Number
	//     pow              	Calculate base exponential
	//       -b, --base     	Base
	//       -e, --exponent 	Exponent

	resetArgs()
}

func ExampleNew_usage_h() {
	os.Args = []string{"gocmd.test", "-h"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
			Help bool `short:"h" long:"help" description:"Display usage" global:"true"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// Usage: basic [options...]
	//
	// A basic app
	//
	// Options:
	//   -h, --help 	Display usage

	resetArgs()
}

func ExampleNew_usage_help() {
	os.Args = []string{"gocmd.test", "--help"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
			Help bool `long:"help" description:"Display usage" global:"true"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// Usage: basic [options...]
	//
	// A basic app
	//
	// Options:
	//       --help 	Display usage

	resetArgs()
}

func ExampleNew_version() {
	os.Args = []string{"gocmd.test", "-vv"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
			Version   bool `short:"v" long:"version" description:"Display version"`
			VersionEx bool `long:"vv" description:"Display version (extended)"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// App name    : basic
	// App version : 1.0.0
	// Go version  : vTest

	resetArgs()
}

func ExampleNew_version_v() {
	os.Args = []string{"gocmd.test", "-v"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
			Version bool `short:"v" long:"version" description:"Display version"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// 1.0.0

	resetArgs()
}

func ExampleNew_version_version() {
	os.Args = []string{"gocmd.test", "--version"}

	gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags: &struct {
			Version bool `long:"version" description:"Display version"`
		}{},
		ConfigType: gocmd.ConfigTypeAuto,
	})
	// Output:
	// 1.0.0

	resetArgs()
}

func ExampleNew_command() {
	os.Args = []string{"gocmd.test", "math", "sqrt", "-n=9"}

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
	// Output:
	// 3

	resetArgs()
}

func ExampleCmd_PrintVersion() {
	cmd, err := gocmd.New(gocmd.Options{
		Version:    "1.0.0",
		ConfigType: gocmd.ConfigTypeAuto,
	})
	if err == nil {
		cmd.PrintVersion(false)
	}
	// Output: 1.0.0
}
