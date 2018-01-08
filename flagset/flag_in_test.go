/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlag_ID(t *testing.T) {
	Convey("should return the id of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ID(), ShouldEqual, 0)
	})
}

func TestFlag_Name(t *testing.T) {
	Convey("should return the name of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Name(), ShouldEqual, "Test")
	})
}

func TestFlag_Short(t *testing.T) {
	Convey("should return the short argument of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Short(), ShouldEqual, "f")
	})
}

func TestFlag_Long(t *testing.T) {
	Convey("should return the long argument of the flag", t, func() {
		flags := struct {
			Test string `long:"foo"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "foo")
	})
}

func TestFlag_Command(t *testing.T) {
	Convey("should return the command argument of the flag", t, func() {
		flags := struct {
			Test struct {
			} `command:"bar"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Command(), ShouldEqual, "bar")
	})
}

func TestFlag_Description(t *testing.T) {
	Convey("should return the description of the flag", t, func() {
		flags := struct {
			Test string `short:"f" description:"baz"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Description(), ShouldEqual, "baz")
	})
}

func TestFlag_Required(t *testing.T) {
	Convey("should return the required value of the flag", t, func() {
		flags := struct {
			Test string `short:"f" required:"true"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Required(), ShouldEqual, true)
	})
}

func TestFlag_Env(t *testing.T) {
	Convey("should return the env value of the flag", t, func() {
		flags := struct {
			Test string `short:"f" env:"GOPATH"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Env(), ShouldEqual, "GOPATH")
	})
}

func TestFlag_Delimiter(t *testing.T) {
	Convey("should return the delimiter value of the flag", t, func() {
		flags := struct {
			Test []string `short:"f" delimiter:","`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Delimiter(), ShouldEqual, ",")
	})
}

func TestFlag_ValueDefault(t *testing.T) {
	Convey("should return the default value of the flag", t, func() {
		flags := struct {
			Test string `short:"f" default:"qux"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueDefault(), ShouldEqual, "qux")
	})
}

func TestFlag_ValueType(t *testing.T) {
	Convey("should return the value type of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueType(), ShouldEqual, "string")
	})
}

func TestFlag_ValueBy(t *testing.T) {
	Convey("should return the value by of the flag", t, func() {
		flags01 := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags01})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "")

		flags02 := struct {
			Test string `short:"f"`
		}{}
		args := []string{"./app", "-f=foo"}
		flagSet, err = New(Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "arg")

		flags03 := struct {
			Test string `short:"f" default:"bar"`
		}{}
		args = []string{"./app"}
		flagSet, err = New(Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "default")

		flags04 := struct {
			Test string `short:"f" env:"FOO_BAR_BAZ"`
		}{}
		args = []string{"./app"}
		flagSet, err = New(Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "")

		flags05 := struct {
			Test string `short:"f" env:"GOPATH"`
		}{}
		args = []string{"./app"}
		flagSet, err = New(Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "env")
	})
}

func TestFlag_Kind(t *testing.T) {
	Convey("should return the kind of the flag", t, func() {
		flags := struct {
			Test    string `short:"t"`
			Command struct {
			}
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Kind(), ShouldEqual, "arg")
		flag = flagSet.FlagByName("Command")
		So(flag, ShouldNotBeNil)
		So(flag.Kind(), ShouldEqual, "command")
	})
}

func TestFlag_ParentID(t *testing.T) {
	Convey("should return the parent id of the flag", t, func() {
		flags := struct {
			CommandFoo struct {
				Test       string `short:"t"`
				CommandBar struct {
					Test       string `short:"t"`
					CommandBaz struct {
						Test string `short:"t"`
					}
				}
			}
		}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("CommandFoo.Test")
		So(flag, ShouldNotBeNil)
		So(flag.ParentID(), ShouldEqual, 0)
		flag = flagSet.FlagByName("CommandFoo.CommandBar.Test")
		So(flag, ShouldNotBeNil)
		So(flag.ParentID(), ShouldEqual, 2)
		flag = flagSet.FlagByName("CommandFoo.CommandBar.CommandBaz.Test")
		So(flag, ShouldNotBeNil)
		So(flag.ParentID(), ShouldEqual, 4)
	})
}
