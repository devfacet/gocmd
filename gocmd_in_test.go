/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocmd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
						Quux bool `short:"q" long:"quux" description:"Test quux"`
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
		So(usageItems, ShouldHaveLength, 4)
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
		So(usage, ShouldEqual, "\nUsage: test [options...] COMMAND [options...]\n\nTest\n\nOptions:       \t\n  -f, --foo    \tTest foo                \t\n  -b           \tTest bar                \t\n      --baz    \tTest baz                \t\n               \t\nCommands:      \t\n  qux          \tQux command             \t\n    -f, --foo  \tTest foo                \t\n        --quux \tTest quux (default test)\t\n")
	})
}
