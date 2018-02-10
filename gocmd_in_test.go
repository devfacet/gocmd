/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocmd

import (
	"os"
	"testing"

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

func TestCmd_usageItems(t *testing.T) {
	Convey("should return the correct usage items", t, func() {
		cmd, err := New(Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"f" long:"foo" description:"Test foo"`
				Bar struct {
					Baz bool `short:"b" long:"baz" description:"Test baz"`
					Qux struct {
						Quux    bool   `short:"q" long:"quux" description:"Test quux"`
						String  string `short:"s" long:"string" default:"/go" env:"GOPATH" description:"Test"`
						Default string `short:"d" long:"default" default:"default" description:"Test"`
						Env     string `short:"e" long:"env" env:"GOPATH" description:"Test"`
					} `description:"Qux command"`
				} `description:"Bar command"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)

		usageItems := cmd.usageItems("arg", -1, 0)
		So(usageItems, ShouldNotBeNil)
		So(usageItems, ShouldHaveLength, 1)
		So(usageItems[0].kind, ShouldEqual, "arg")
		So(usageItems[0].flagID, ShouldEqual, 0)
		So(usageItems[0].parentID, ShouldEqual, -1)
		So(usageItems[0].left, ShouldEqual, "-f, --foo")
		So(usageItems[0].right, ShouldEqual, "Test foo")
		So(usageItems[0].level, ShouldEqual, 1)

		usageItems = cmd.usageItems("command", -1, 0)
		So(usageItems, ShouldNotBeNil)
		So(usageItems, ShouldHaveLength, 7)
		So(usageItems[0].kind, ShouldEqual, "command")
		So(usageItems[0].flagID, ShouldEqual, 1)
		So(usageItems[0].parentID, ShouldEqual, -1)
		So(usageItems[0].left, ShouldEqual, "bar")
		So(usageItems[0].right, ShouldEqual, "Bar command")
		So(usageItems[0].level, ShouldEqual, 1)
		So(usageItems[1].kind, ShouldEqual, "arg")
		So(usageItems[1].flagID, ShouldEqual, 2)
		So(usageItems[1].parentID, ShouldEqual, 1)
		So(usageItems[1].left, ShouldEqual, "-b, --baz")
		So(usageItems[1].right, ShouldEqual, "Test baz")
		So(usageItems[1].level, ShouldEqual, 2)
		So(usageItems[2].kind, ShouldEqual, "command")
		So(usageItems[2].flagID, ShouldEqual, 3)
		So(usageItems[2].parentID, ShouldEqual, 1)
		So(usageItems[2].left, ShouldEqual, "qux")
		So(usageItems[2].right, ShouldEqual, "Qux command")
		So(usageItems[2].level, ShouldEqual, 2)
		So(usageItems[3].kind, ShouldEqual, "arg")
		So(usageItems[3].flagID, ShouldEqual, 4)
		So(usageItems[3].parentID, ShouldEqual, 3)
		So(usageItems[3].left, ShouldEqual, "-q, --quux")
		So(usageItems[3].right, ShouldEqual, "Test quux")
		So(usageItems[3].level, ShouldEqual, 3)
		So(usageItems[4].left, ShouldEqual, "-s, --string")
		So(usageItems[4].right, ShouldEqual, "Test (default /go - override $GOPATH)")
		So(usageItems[4].level, ShouldEqual, 3)
		So(usageItems[5].left, ShouldEqual, "-d, --default")
		So(usageItems[5].right, ShouldEqual, "Test (default default)")
		So(usageItems[5].level, ShouldEqual, 3)
		So(usageItems[6].left, ShouldEqual, "-e, --env")
		So(usageItems[6].right, ShouldEqual, "Test (default $GOPATH)")
		So(usageItems[6].level, ShouldEqual, 3)
	})
}

func TestCmd_usageContent(t *testing.T) {
	Convey("should return correct usage content", t, func() {
		cmd, err := New(Options{
			Name:        "test",
			Version:     "1.0.0",
			Description: "Test",
			Flags: &struct {
				Foo bool `short:"f" long:"foo" description:"Test foo"`
				Bar bool `short:"b" description:"Test bar"`
				Baz bool `long:"baz" description:"Test baz"`
				Qux struct {
					Foo  bool `short:"f" long:"foo" description:"Test foo"`
					Quux bool `long:"quux" default:"test" description:"Test quux"`
				} `description:"Qux command"`
			}{},
		})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		usage := cmd.usageContent()
		So(usage, ShouldNotBeEmpty)
		So(usage, ShouldEqual, "\nUsage: test [options...] COMMAND [options...]\n\nTest\n\nOptions:\n  -f, --foo    \tTest foo\n  -b           \tTest bar\n      --baz    \tTest baz\n\nCommands:\n  qux          \tQux command\n    -f, --foo  \tTest foo\n        --quux \tTest quux (default test)\n")
	})
}

func TestCmd_isTest(t *testing.T) {
	Convey("should return whether it's a test or not", t, func() {
		cmd, err := New(Options{Name: "test"})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.isTest(), ShouldEqual, true)

		os.Args = []string{"./app", "-test."}
		cmd, err = New(Options{Name: "test"})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.isTest(), ShouldEqual, true)
		resetArgs()

		os.Args = []string{"./app"}
		cmd, err = New(Options{Name: "test"})
		So(err, ShouldBeNil)
		So(cmd, ShouldNotBeNil)
		So(cmd.isTest(), ShouldEqual, false)
		resetArgs()
	})
}
