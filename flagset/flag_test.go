/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset_test

import (
	"errors"
	"testing"

	"github.com/devfacet/gocmd/flagset"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFlag_ID(t *testing.T) {
	Convey("should return the id of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "foo")
	})
}

func TestFlag_FormattedArg(t *testing.T) {
	Convey("should return the formatted argument of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Short(), ShouldEqual, "f")
	})
}

func TestFlag_Command(t *testing.T) {
	Convey("should return the command argument of the flag", t, func() {
		flags := struct {
			Test struct {
			} `command:"bar"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Required(), ShouldEqual, true)
	})
}

func TestFlag_Global(t *testing.T) {
	Convey("should return the global value of the flag", t, func() {
		flags := struct {
			Test string `short:"f" global:"true"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Global(), ShouldEqual, true)
	})
}

func TestFlag_Env(t *testing.T) {
	Convey("should return the env value of the flag", t, func() {
		flags := struct {
			Test string `short:"f" env:"GOPATH"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "")

		flags02 := struct {
			Test string `short:"f"`
		}{}
		args := []string{"./app", "-f=foo"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "arg")

		flags03 := struct {
			Test string `short:"f" default:"bar"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "default")

		flags04 := struct {
			Test string `short:"f" env:"FOO_BAR_BAZ"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "")

		flags05 := struct {
			Test string `short:"f" env:"GOPATH"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag = flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.ValueBy(), ShouldEqual, "env")
	})
}

func TestFlag_Value(t *testing.T) {
	Convey("should return the value of the flag", t, func() {
		flags := struct {
			Test string `short:"f"`
		}{}
		args := []string{"./app", "-f=foo"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Value(), ShouldEqual, "foo")
	})
}

func TestFlag_Kind(t *testing.T) {
	Convey("should return the kind of the flag", t, func() {
		flags := struct {
			Test    string `short:"t"`
			Command struct {
			}
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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

func TestFlag_FieldIndex(t *testing.T) {
	Convey("should return the field index of the flag", t, func() {
		flags := struct {
			Foo string `short:"f"`
			Bar struct {
				Baz string `short:"b"`
			} `command:"bar"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Foo")
		So(flag, ShouldNotBeNil)
		So(flag.FieldIndex(), ShouldHaveLength, 1)
		So(flag.FieldIndex()[0], ShouldEqual, 0)
		flag = flagSet.FlagByName("Bar")
		So(flag, ShouldNotBeNil)
		So(flag.FieldIndex(), ShouldHaveLength, 1)
		So(flag.FieldIndex(), ShouldResemble, []int{1})
		flag = flagSet.FlagByName("Bar.Baz")
		So(flag, ShouldNotBeNil)
		So(flag.FieldIndex(), ShouldHaveLength, 2)
		So(flag.FieldIndex(), ShouldResemble, []int{1, 0})
	})
}

func TestFlag_ParentIndex(t *testing.T) {
	Convey("should return the field index of the flag", t, func() {
		flags := struct {
			Foo string `short:"f"`
			Bar struct {
				Baz string `short:"b"`
				Qux struct {
					Quux string `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Bar.Baz")
		So(flag, ShouldNotBeNil)
		So(flag.ParentIndex(), ShouldHaveLength, 1)
		So(flag.ParentIndex(), ShouldResemble, []int{1})
		flag = flagSet.FlagByName("Bar.Qux.Quux")
		So(flag, ShouldNotBeNil)
		So(flag.ParentIndex(), ShouldHaveLength, 2)
		So(flag.ParentIndex(), ShouldResemble, []int{1, 1})
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
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

func TestFlag_CommandID(t *testing.T) {
	Convey("should return the command id of the flag", t, func() {
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
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("CommandFoo")
		So(flag, ShouldNotBeNil)
		So(flag.CommandID(), ShouldEqual, 0)
		flag = flagSet.FlagByName("CommandFoo.CommandBar")
		So(flag, ShouldNotBeNil)
		So(flag.CommandID(), ShouldEqual, 1)
		flag = flagSet.FlagByName("CommandFoo.CommandBar.CommandBaz")
		So(flag, ShouldNotBeNil)
		So(flag.CommandID(), ShouldEqual, 2)
	})
}

func TestFlag_Err(t *testing.T) {
	Convey("should return the error of the flag", t, func() {
		flags := struct {
			Test string `short:"f" required:"true"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flag := flagSet.FlagByName("Test")
		So(flag, ShouldNotBeNil)
		So(flag.Err(), ShouldBeError, errors.New("argument -f is required"))
	})
}
