/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocmd_test

import (
	"errors"
	"os"
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
		// Override the command line arguments for the test
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
	})

	Convey("should fail to create a new command", t, func() {
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
	})
}

func TestCmd_FlagErrors(t *testing.T) {
	Convey("should return the flag errors", t, func() {
		cmd, err := gocmd.New(gocmd.Options{
			Flags: &struct {
				Foo struct {
				} `global:"true"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		flagErrors := cmd.FlagErrors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("command foo can't be global"))
	})
}

func ExampleCmd_PrintVersion() {
	cmd, _ := gocmd.New(gocmd.Options{
		Version: "1.0.0",
	})
	cmd.PrintVersion(false)
	// Output: 1.0.0
}

func ExampleCmd_PrintVersion_extra() {
	cmd, _ := gocmd.New(gocmd.Options{
		Name:    "test",
		Version: "1.0.0",
	})
	cmd.PrintVersion(true)
	// Output:
	// App name    : test
	// App version : 1.0.0
	// Go version  : TEST
}

func ExampleCmd_PrintUsage() {
	cmd, _ := gocmd.New(gocmd.Options{
		Name:        "test",
		Version:     "1.0.0",
		Description: "Test",
	})
	cmd.PrintUsage()
	// Output:
	// Usage: test
	//
	// Test
}
