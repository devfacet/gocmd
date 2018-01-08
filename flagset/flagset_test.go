/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/devfacet/gocmd/flagset"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("should fail to create a new flag set", t, func() {
		flagSet, err := flagset.New(flagset.Options{})
		So(err, ShouldBeError, errors.New("flags are required"))
		So(flagSet, ShouldBeNil)

		flagSet, err = flagset.New(flagset.Options{Flags: ""})
		So(err, ShouldBeError, errors.New("flags must be a struct pointer"))
		So(flagSet, ShouldBeNil)

		s := ""
		flagSet, err = flagset.New(flagset.Options{Flags: &s})
		So(err, ShouldBeError, errors.New("flags must be a struct pointer"))
		So(flagSet, ShouldBeNil)

		flagSet, err = flagset.New(flagset.Options{Flags: struct{}{}})
		So(err, ShouldBeError, errors.New("flags must be a struct pointer"))
		So(flagSet, ShouldBeNil)

		var i interface{}
		flagSet, err = flagset.New(flagset.Options{Flags: &i})
		So(err, ShouldBeError, errors.New("flags must be a struct pointer"))
		So(flagSet, ShouldBeNil)

		flags01 := struct {
			Foo []*string `long:"foo"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags01})
		So(err, ShouldBeError, errors.New("invalid type []*string. Supported types: [bool float64 int int64 uint uint64 string []bool []float64 []int []int64 []uint []uint64 []string struct]"))
		So(flagSet, ShouldBeNil)

		flags02 := struct {
			Version bool `short:"v"`
			Verbose bool `short:"v"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02})
		So(err, ShouldBeError, errors.New("short argument v in Verbose field is already defined in Version field"))
		So(flagSet, ShouldBeNil)

		flags03 := struct {
			Log     bool `long:"verbose"`
			Verbose bool `long:"verbose"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03})
		So(err, ShouldBeError, errors.New("long argument verbose in Verbose field is already defined in Log field"))
		So(flagSet, ShouldBeNil)

		flags04 := struct {
			Foo struct{} `command:"foo"`
			Bar struct{} `command:"foo"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04})
		So(err, ShouldBeError, errors.New("command foo in Bar field is already defined in Foo field"))
		So(flagSet, ShouldBeNil)

		flags05 := struct {
			Foo struct {
				Foo bool `short:"f" long:"foo"`
				Bar bool `short:"f" long:"foo"`
			}
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05})
		So(err, ShouldBeError, errors.New("short argument f in Bar field is already defined in Foo field"))
		So(flagSet, ShouldBeNil)

		flags06 := struct {
			Foo struct {
				Foo bool `short:"f" long:"foo"`
				Bar bool `short:"f" long:"foo"`
			} `command:"foo"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06})
		So(err, ShouldBeError, errors.New("short argument f in Bar field is already defined in Foo field"))
		So(flagSet, ShouldBeNil)

		flags07 := struct {
			Version bool `short:"vv"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags07})
		So(err, ShouldBeError, errors.New("short argument vv in Version field must be one character long"))
		So(flagSet, ShouldBeNil)
	})

	Convey("should return a new flag set", t, func() {
		s := struct{}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &s})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
	})

	Convey("should return correct errors", t, func() {
		flags01 := struct {
			Bool    bool      `short:"b" long:"bool"`
			Float64 float64   `short:"f" long:"float64"`
			Int     int       `short:"i" long:"int"`
			Int64   int64     `short:"I" long:"int64"`
			Uint    uint      `short:"u" long:"uint"`
			Uint64  uint64    `short:"U" long:"uint64"`
			Bools   []bool    `long:"bools" delimiter:","`
			Floats  []float64 `long:"floats" delimiter:","`
			Ints    []int     `long:"ints" delimiter:","`
			Int64s  []int64   `long:"Ints" delimiter:","`
			Uints   []uint    `long:"uints" delimiter:","`
			Uint64s []uint64  `long:"Uints" delimiter:","`
			Env     int       `short:"e" env:"GOPATH"`
			Default int       `short:"d" default:"DEFAULT"`
		}{}
		args := []string{
			"./app",
			"-b=foo",
			"-f=foo",
			"-i=foo",
			"-I=foo",
			"-u=foo",
			"-U=foo",
			"--bools=true,foofoo,false",
			"--floats=0.1,foofoo,0.2",
			"--ints=1,foofoo,2",
			"--Ints=1,foofoo,2",
			"--uints=1,foofoo,2",
			"--Uints=1,foofoo,2",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as bool"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as float64"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as int"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as int64"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as uint"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as uint64"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foofoo' as bool"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foofoo' as float64"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foofoo' as int"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foofoo' as int64"))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foofoo' as uint"))
		So(flagErrors, ShouldContain, fmt.Errorf("failed to parse '%s' as int", os.Getenv("GOPATH")))
		So(flagErrors, ShouldContain, errors.New("failed to parse 'DEFAULT' as int"))
	})

	Convey("should return correct flags (sanity)", t, func() {
		flags01 := struct {
			Foo        bool `short:" "`
			Bar        bool `long:" "`
			Baz        bool `short:" f "`
			Qux        bool `long:" qux "`
			CommandFoo struct {
				Foo bool `short:"f"`
			} `command:" foo "`
		}{}
		args := []string{
			"./app",
			"-f",
			"--qux",
			"foo",
			"-f",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Foo, ShouldEqual, false)
		So(flags01.Bar, ShouldEqual, false)
		So(flags01.Baz, ShouldEqual, true)
		So(flags01.Qux, ShouldEqual, true)
		So(flags01.CommandFoo.Foo, ShouldEqual, true)

		flags02 := struct {
			Foo        bool `short:"\"f\""`
			Bar        bool `long:"\"bar\""`
			CommandFoo struct {
				Foo bool `short:"f"`
			} `command:"\"foo\""`
		}{}
		args = []string{
			"./app",
			"-f",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Foo, ShouldEqual, true)
		So(flags02.Bar, ShouldEqual, false)
		So(flags02.CommandFoo.Foo, ShouldEqual, false)

		flags03 := struct {
			CommandFoo struct {
				Foo bool `short:"f"`
			} `command:" foo "`
		}{}
		args = []string{
			"./app",
			"foo",
			"-f",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.CommandFoo.Foo, ShouldEqual, true)

		flags04 := struct {
			Foo        string `long:"foo-foo"`
			Bar        string `long:"bar.bar"`
			Baz        string `long:"baz_baz"`
			CommandFoo struct {
				Foo        string `long:"foo-foo"`
				Bar        string `long:"bar.bar"`
				Baz        string `long:"baz_baz"`
				CommandBar struct {
					Test string `long:"test"`
				} `command:"command.bar"`
			} `command:"command-foo"`
			CommandBaz struct{} `command:"command_baz"`
			Qux        string   `long:"qux"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flags := flagSet.Flags()
		So(flags[0].Long(), ShouldEqual, "foo-foo")
		So(flags[1].Long(), ShouldEqual, "bar.bar")
		So(flags[2].Long(), ShouldEqual, "baz_baz")
		So(flags[3].Command(), ShouldEqual, "command-foo")
		So(flags[4].Long(), ShouldEqual, "foo-foo")
		So(flags[5].Long(), ShouldEqual, "bar.bar")
		So(flags[6].Long(), ShouldEqual, "baz_baz")
		So(flags[7].Command(), ShouldEqual, "command.bar")
		So(flags[8].Long(), ShouldEqual, "test")
		So(flags[9].Command(), ShouldEqual, "command_baz")
		So(flags[10].Long(), ShouldEqual, "qux")
	})

	Convey("should return correct flag values (last)", t, func() {
		flags01 := struct {
			String string `short:"s" long:"string" default:"foo"`
		}{}
		args := []string{"./app", "-s=bar", "--string=baz"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.String, ShouldEqual, "baz")

		flags02 := struct {
			String string `short:"s" long:"string" default:"foo"`
		}{}
		args = []string{"./app", "--string=baz", "-s=bar"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.String, ShouldEqual, "bar")

		flags03 := struct {
			String string `short:"s" long:"string" default:"foo"`
		}{}
		args = []string{"./app", "-s=bar", "-s=baz", "-s=qux"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.String, ShouldEqual, "qux")

		flags04 := struct {
			Bool bool `short:"b" long:"bool" default:"true"`
		}{}
		args = []string{"./app", "-b=true", "-b=true", "-b=false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.Bool, ShouldEqual, false)

		flags05 := struct {
			Bool bool `short:"b" long:"bool"`
		}{}
		args = []string{"./app", "-b=true", "-b=true", "-b=false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.Bool, ShouldEqual, false)

		flags06 := struct {
			Bool bool `short:"b" long:"bool" default:"true"`
		}{}
		args = []string{"./app", "-b=false", "-b=false", "-b=true"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags06.Bool, ShouldEqual, true)
	})

	Convey("should return correct flag values (default)", t, func() {
		flags01 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args := []string{"./app"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Default, ShouldEqual, "foo")

		flags02 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d=bar"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Default, ShouldEqual, "bar")

		flags03 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d=\"bar\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Default, ShouldEqual, "bar")

		flags04 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d", "bar"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.Default, ShouldEqual, "bar")

		flags05 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.Default, ShouldEqual, "")

		flags06 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags06.Default, ShouldEqual, "")

		flags07 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d", "\"foo bar\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags07.Default, ShouldEqual, "foo bar")

		flags10 := struct {
			Default string `short:"d" long:"default" default:"foo"`
		}{}
		args = []string{"./app", "-d"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags10, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -d needs a value"))
		So(flags10.Default, ShouldEqual, "")
	})

	Convey("should return correct flag values (required)", t, func() {
		flags01 := struct {
			Foo    bool   `short:"f" long:"foo" required:"true"`
			String string `short:"s" long:"string" required:"true"`
		}{}
		args := []string{"./app"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f is required"))
		So(flagErrors, ShouldContain, errors.New("argument -s is required"))
		So(flags01.Foo, ShouldEqual, false)
		So(flags01.String, ShouldEqual, "")

		flags02 := struct {
			Foo    bool   `long:"foo" required:"true"`
			String string `long:"string" required:"true"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument --foo is required"))
		So(flagErrors, ShouldContain, errors.New("argument --string is required"))
		So(flags02.Foo, ShouldEqual, false)
		So(flags02.String, ShouldEqual, "")

		flags03 := struct {
			CommandFoo struct {
				Foo bool `short:"f" long:"foo" required:"true"`
			} `command:"foo"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.CommandFoo.Foo, ShouldEqual, false)

		flags04 := struct {
			CommandFoo struct {
				Foo bool `short:"f" long:"foo" required:"true"`
			} `command:"bar" required:"true"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("command bar is required"))
		So(flags04.CommandFoo.Foo, ShouldEqual, false)

		flags05 := struct {
			CommandFoo struct {
				Foo    bool   `short:"f" long:"foo" required:"true"`
				String string `long:"string" required:"true"`
			} `command:"bar" required:"true"`
		}{}
		args = []string{
			"./app",
			"bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f is required for bar command"))
		So(flags05.CommandFoo.Foo, ShouldEqual, false)
		So(flags05.CommandFoo.String, ShouldEqual, "")

		flags06 := struct {
			Foo        bool   `short:"f" long:"foo" required:"true"`
			String     string `short:"s" long:"string" required:"true"`
			CommandQux struct {
				Baz bool `short:"b" long:"baz" required:"true"`
			} `command:"qux" required:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s=foo",
			"qux",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags06.Foo, ShouldEqual, true)
		So(flags06.String, ShouldEqual, "foo")
		So(flags06.CommandQux.Baz, ShouldEqual, true)
	})

	Convey("should return correct flag values (env)", t, func() {
		flags01 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args := []string{"./app"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Env, ShouldEqual, os.Getenv("GOPATH"))

		flags02 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args = []string{"./app", "-e=/go"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Env, ShouldEqual, "/go")

		flags03 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args = []string{"./app", "-e="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Env, ShouldEqual, "")

		flags04 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args = []string{"./app", "-e=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.Env, ShouldEqual, "")

		flags05 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args = []string{"./app", "-e=''"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.Env, ShouldEqual, "")

		flags10 := struct {
			Env string `short:"e" long:"env" env:"GOPATH"`
		}{}
		args = []string{"./app", "-e"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags10, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -e needs a value"))
		So(flags10.Env, ShouldEqual, "")
	})

	Convey("should return correct flag values (bool)", t, func() {
		flags01 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args := []string{"./app"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Foo, ShouldEqual, false)

		flags02 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Foo, ShouldEqual, true)

		flags03 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f=true"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Foo, ShouldEqual, true)

		flags04 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f=\"true\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.Foo, ShouldEqual, true)

		flags05 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f", "true"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.Foo, ShouldEqual, true)

		flags06 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f", "\"true\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags06.Foo, ShouldEqual, true)

		flags07 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f=false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags07.Foo, ShouldEqual, false)

		flags08 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f=\"false\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags08, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags08.Foo, ShouldEqual, false)

		flags09 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f", "false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags09, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags09.Foo, ShouldEqual, false)

		flags10 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f", "\"false\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags10, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags10.Foo, ShouldEqual, false)

		flags11 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags11, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags11.Foo, ShouldEqual, false)

		flags12 := struct {
			Foo bool `short:"f" long:"foo"`
		}{}
		args = []string{"./app", "-f=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags12, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags12.Foo, ShouldEqual, false)

		flags13 := struct {
			Foo bool `short:"f" long:"foo" default:"true"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags13, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags13.Foo, ShouldEqual, true)

		flags14 := struct {
			Foo bool `short:"f" long:"foo" default:"false"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags14, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags14.Foo, ShouldEqual, false)

		flags15 := struct {
			Foo bool `short:"f" long:"foo" default:"false"`
		}{}
		args = []string{"./app", "-f"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags15, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags15.Foo, ShouldEqual, true)

		flags16 := struct {
			Foo bool `short:"f" long:"foo" default:"true"`
		}{}
		args = []string{"./app", "-f", "false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags16, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags16.Foo, ShouldEqual, false)

		flags17 := struct {
			Foo bool `short:"f" long:"foo" default:"true"`
		}{}
		args = []string{"./app", "-f="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags17, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags17.Foo, ShouldEqual, false)

		flags18 := struct {
			Foo bool `short:"f" long:"foo" default:"true"`
		}{}
		args = []string{"./app", "-f=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags18, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags18.Foo, ShouldEqual, false)

		flags19 := struct {
			Foo bool `short:"f" long:"foo" default:"false"`
		}{}
		args = []string{"./app", "-f="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags19, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags19.Foo, ShouldEqual, false)

		flags20 := struct {
			Foo bool `short:"f" long:"foo" default:"false"`
		}{}
		args = []string{"./app", "-f=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags20, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags20.Foo, ShouldEqual, false)
	})

	Convey("should return correct flag values (float)", t, func() {
		flags01 := struct {
			Float float64 `short:"f" long:"float"`
		}{}
		args := []string{"./app", "-f"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags01.Float, ShouldEqual, 0)

		flags02 := struct {
			Float float64 `short:"f" long:"float"`
		}{}
		args = []string{"./app", "-f", "0.2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Float, ShouldEqual, 0.2)

		flags03 := struct {
			Float float64 `short:"f" long:"float" default:"0.3"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Float, ShouldEqual, 0.3)

		flags04 := struct {
			Float float64 `short:"f" long:"float" default:"0.4"`
		}{}
		args = []string{"./app", "-f="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags04.Float, ShouldEqual, 0)

		flags05 := struct {
			Float float64 `short:"f" long:"float" default:"0.5"`
		}{}
		args = []string{"./app", "-f=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags05.Float, ShouldEqual, 0)
	})

	Convey("should return correct flag values (int)", t, func() {
		flags01 := struct {
			Int int `short:"i" long:"int"`
		}{}
		args := []string{"./app", "-i"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags01.Int, ShouldEqual, 0)

		flags02 := struct {
			Int int `short:"i" long:"int"`
		}{}
		args = []string{"./app", "-i", "2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Int, ShouldEqual, 2)

		flags03 := struct {
			Int int `short:"i" long:"int" default:"3"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Int, ShouldEqual, 3)

		flags04 := struct {
			Int int `short:"i" long:"int" default:"4"`
		}{}
		args = []string{"./app", "-i="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags04.Int, ShouldEqual, 0)

		flags05 := struct {
			Int int `short:"i" long:"int" default:"5"`
		}{}
		args = []string{"./app", "-i=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags05.Int, ShouldEqual, 0)
	})

	Convey("should return correct flag values (int64)", t, func() {
		flags01 := struct {
			Int64 int64 `short:"i" long:"int64"`
		}{}
		args := []string{"./app", "-i"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags01.Int64, ShouldEqual, 0)

		flags02 := struct {
			Int64 int64 `short:"i" long:"int64"`
		}{}
		args = []string{"./app", "-i", "2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Int64, ShouldEqual, 2)

		flags03 := struct {
			Int64 int64 `short:"i" long:"int64" default:"3"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Int64, ShouldEqual, 3)

		flags04 := struct {
			Int64 int64 `short:"i" long:"int64" default:"4"`
		}{}
		args = []string{"./app", "-i="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags04.Int64, ShouldEqual, 0)

		flags05 := struct {
			Int64 int64 `short:"i" long:"int64" default:"4"`
		}{}
		args = []string{"./app", "-i=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags05.Int64, ShouldEqual, 0)
	})

	Convey("should return correct flag values (uint)", t, func() {
		flags01 := struct {
			Uint uint `short:"u" long:"uint"`
		}{}
		args := []string{"./app", "-u"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags01.Uint, ShouldEqual, 0)

		flags02 := struct {
			Uint uint `short:"u" long:"uint"`
		}{}
		args = []string{"./app", "-u", "2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Uint, ShouldEqual, 2)

		flags03 := struct {
			Uint uint `short:"u" long:"uint" default:"3"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Uint, ShouldEqual, 3)

		flags04 := struct {
			Uint uint `short:"u" long:"uint" default:"4"`
		}{}
		args = []string{"./app", "-u="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags04.Uint, ShouldEqual, 0)

		flags05 := struct {
			Uint uint `short:"u" long:"uint" default:"5"`
		}{}
		args = []string{"./app", "-u=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags05.Uint, ShouldEqual, 0)
	})

	Convey("should return correct flag values (uint64)", t, func() {
		flags01 := struct {
			Uint64 uint64 `short:"u" long:"uint64"`
		}{}
		args := []string{"./app", "-u"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags01.Uint64, ShouldEqual, 0)

		flags02 := struct {
			Uint64 uint64 `short:"u" long:"uint64"`
		}{}
		args = []string{"./app", "-u", "2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Uint64, ShouldEqual, 2)

		flags03 := struct {
			Uint64 uint64 `short:"u" long:"uint64" default:"3"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Uint64, ShouldEqual, 3)

		flags04 := struct {
			Uint64 uint64 `short:"u" long:"uint64" default:"4"`
		}{}
		args = []string{"./app", "-u=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags04.Uint64, ShouldEqual, 0)

		flags05 := struct {
			Uint64 uint64 `short:"u" long:"uint64" default:"5"`
		}{}
		args = []string{"./app", "-u=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags05.Uint64, ShouldEqual, 0)
	})

	Convey("should return correct flag values (string)", t, func() {
		flags01 := struct {
			String string `short:"s" long:"string"`
		}{}
		args := []string{"./app", "-s"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -s needs a value"))
		So(flags01.String, ShouldEqual, "")

		flags02 := struct {
			String string `short:"s" long:"string"`
		}{}
		args = []string{"./app", "-s", "foo"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.String, ShouldEqual, "foo")

		flags03 := struct {
			String string `short:"s" long:"string" default:"bar"`
		}{}
		args = []string{"./app"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.String, ShouldEqual, "bar")

		flags04 := struct {
			String string `short:"s" long:"string" default:"baz"`
		}{}
		args = []string{"./app", "-s="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.String, ShouldEqual, "")

		flags05 := struct {
			String string `short:"s" long:"string" default:"baz"`
		}{}
		args = []string{"./app", "-s=\"\""}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.String, ShouldEqual, "")
	})

	Convey("should return correct flag values ([]bool)", t, func() {
		flags01 := struct {
			Bools []bool `short:"b" long:"bools"`
		}{}
		args := []string{"./app", "-b=true", "--bools=true", "-b=false"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Bools, ShouldResemble, []bool{true, true, false})

		flags02 := struct {
			Bools []bool `short:"b" long:"bools"`
		}{}
		args = []string{"./app", "-b", "--bools=true", "-b=false"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Bools, ShouldResemble, []bool{true, true, false})

		flags03 := struct {
			Bools []bool `short:"b"`
		}{}
		args = []string{"./app", "-b", "-b="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -b needs a value"))
		So(flags03.Bools, ShouldBeNil)
	})

	Convey("should return correct flag values ([]float)", t, func() {
		flags01 := struct {
			Floats []float64 `short:"f" long:"floats"`
		}{}
		args := []string{"./app", "-f=0.1", "--floats=0.2"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Floats, ShouldResemble, []float64{0.1, 0.2})

		flags02 := struct {
			Floats []float64 `short:"f" long:"floats"`
		}{}
		args = []string{"./app", "-f", "-f=0.2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags02.Floats, ShouldBeNil)

		flags03 := struct {
			Floats []float64 `short:"f" long:"floats"`
		}{}
		args = []string{"./app", "-f=0.1", "-f="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -f needs a value"))
		So(flags03.Floats, ShouldBeNil)
	})

	Convey("should return correct flag values ([]int)", t, func() {
		flags01 := struct {
			Ints []int `short:"i" long:"ints"`
		}{}
		args := []string{"./app", "-i=1", "--ints=2"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Ints, ShouldResemble, []int{1, 2})

		flags02 := struct {
			Ints []int `short:"i" long:"ints"`
		}{}
		args = []string{"./app", "-i", "-i=2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags02.Ints, ShouldBeNil)

		flags03 := struct {
			Ints []int `short:"i" long:"ints"`
		}{}
		args = []string{"./app", "-i=1", "-i="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags03.Ints, ShouldBeNil)
	})

	Convey("should return correct flag values ([]int64)", t, func() {
		flags01 := struct {
			Int64s []int64 `short:"i" long:"ints"`
		}{}
		args := []string{"./app", "-i=1", "--ints=2"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Int64s, ShouldResemble, []int64{1, 2})

		flags02 := struct {
			Int64s []int64 `short:"i" long:"ints"`
		}{}
		args = []string{"./app", "-i", "-i=2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags02.Int64s, ShouldBeNil)

		flags03 := struct {
			Int64s []int64 `short:"i" long:"ints"`
		}{}
		args = []string{"./app", "-i=1", "-i="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -i needs a value"))
		So(flags03.Int64s, ShouldBeNil)
	})

	Convey("should return correct flag values ([]uint)", t, func() {
		flags01 := struct {
			Uints []uint `short:"u" long:"uints"`
		}{}
		args := []string{"./app", "-u=1", "--uints=2"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Uints, ShouldResemble, []uint{1, 2})

		flags02 := struct {
			Uints []uint `short:"u" long:"uints"`
		}{}
		args = []string{"./app", "-u", "-u=2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags02.Uints, ShouldBeNil)

		flags03 := struct {
			Uints []uint `short:"u" long:"uints"`
		}{}
		args = []string{"./app", "-u=1", "-u="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags03.Uints, ShouldBeNil)
	})

	Convey("should return correct flag values ([]uint64)", t, func() {
		flags01 := struct {
			Uint64s []uint64 `short:"u" long:"uints"`
		}{}
		args := []string{"./app", "-u=1", "--uints=2"}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Uint64s, ShouldResemble, []uint64{1, 2})

		flags02 := struct {
			Uint64s []uint64 `short:"u" long:"uints"`
		}{}
		args = []string{"./app", "-u", "-u=2"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags02.Uint64s, ShouldBeNil)

		flags03 := struct {
			Uint64s []uint64 `short:"u" long:"uints"`
		}{}
		args = []string{"./app", "-u=1", "-u="}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -u needs a value"))
		So(flags03.Uint64s, ShouldBeNil)
	})

	Convey("should return correct flag values ([]string)", t, func() {
		flags01 := struct {
			Strings []string `short:"s" long:"strings"`
		}{}
		args := []string{"./app", "-s=foo", "--strings=bar", "-s", "baz", "-s=\"foo bar\""}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Strings, ShouldResemble, []string{"foo", "bar", "baz", "foo bar"})

		flags02 := struct {
			Strings []string `short:"s" long:"strings"`
		}{}
		args = []string{"./app", "-s=foo", "-s"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -s needs a value"))
		So(flags02.Strings, ShouldBeNil)

		flags03 := struct {
			Strings []string `short:"s" long:"strings"`
		}{}
		args = []string{"./app", "-s=foo", "-s"}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors = flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -s needs a value"))
		So(flags03.Strings, ShouldBeNil)
	})

	Convey("should return correct flag values (delimiter)", t, func() {
		flags01 := struct {
			Bools   []bool    `short:"b" long:"bools" delimiter:","`
			Floats  []float64 `short:"f" long:"floats" delimiter:","`
			Ints    []int     `short:"i" long:"ints" delimiter:","`
			Int64s  []int64   `short:"I" long:"Ints" delimiter:","`
			Uints   []uint    `short:"u" long:"uints" delimiter:","`
			Uint64s []uint64  `short:"U" long:"Uints" delimiter:","`
			Strings []string  `short:"s" long:"strings" delimiter:","`
		}{}
		args := []string{
			"./app",
			"-b=true,false", "--bools=false,true",
			"-f=0.1,0.2", "--floats=0.3,0.4",
			"-i=1,2", "--ints=3,4",
			"-I=1,2", "--Ints=3,4",
			"-u=1,2", "--uints=3,4",
			"-U=1,2", "--Uints=3,4",
			"-s=foo,bar", "--strings=baz,qux",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Bools, ShouldResemble, []bool{true, false, false, true})
		So(flags01.Floats, ShouldResemble, []float64{0.1, 0.2, 0.3, 0.4})
		So(flags01.Ints, ShouldResemble, []int{1, 2, 3, 4})
		So(flags01.Int64s, ShouldResemble, []int64{1, 2, 3, 4})
		So(flags01.Uints, ShouldResemble, []uint{1, 2, 3, 4})
		So(flags01.Uint64s, ShouldResemble, []uint64{1, 2, 3, 4})
		So(flags01.Strings, ShouldResemble, []string{"foo", "bar", "baz", "qux"})

		flags02 := struct {
			Bools   []bool   `short:"b" long:"bools" delimiter:","`
			Ints    []int    `short:"i" long:"ints" delimiter:","`
			Strings []string `short:"s" long:"strings" delimiter:","`
		}{}
		args = []string{
			"./app",
			"-b=true,,false",
			"-i=1,,",
			"-s=,,,bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Bools, ShouldResemble, []bool{true, false})
		So(flags02.Ints, ShouldResemble, []int{1})
		So(flags02.Strings, ShouldResemble, []string{"bar"})
	})

	Convey("should return correct flag values (command)", t, func() {
		flags01 := struct {
			Foo        bool `short:"f" long:"foo"`
			CommandFoo struct {
				Foo    bool   `short:"f" long:"foo"`
				String string `short:"s" long:"string"`
			} `command:"foo"`
		}{}
		args := []string{
			"./app",
			"-f",
			"-s=foo",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags01.Foo, ShouldEqual, true)
		So(flags01.CommandFoo.Foo, ShouldEqual, false)
		So(flags01.CommandFoo.String, ShouldEqual, "")

		flags02 := struct {
			Foo        bool `short:"f" long:"foo"`
			CommandFoo struct {
				Foo    bool   `short:"f" long:"foo"`
				String string `short:"s" long:"string"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s=bar",
			"foo",
			"-f=false",
			"-s=baz",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Foo, ShouldEqual, true)
		So(flags02.CommandFoo.Foo, ShouldEqual, false)
		So(flags02.CommandFoo.String, ShouldEqual, "baz")

		flags03 := struct {
			Foo        bool `short:"f" long:"foo"`
			CommandFoo struct {
				Foo    bool   `short:"f" long:"foo"`
				String string `short:"s" long:"string"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-f=false",
			"-s=bar",
			"foo",
			"-f",
			"-s=baz",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags03.Foo, ShouldEqual, false)
		So(flags03.CommandFoo.Foo, ShouldEqual, true)
		So(flags03.CommandFoo.String, ShouldEqual, "baz")

		flags04 := struct {
			Foo        bool `short:"f" long:"foo"`
			CommandFoo struct {
				Foo    bool   `short:"f" long:"foo"`
				String string `short:"s" long:"string"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s=bar",
			"foo",
			"-f",
			"-s=baz",
			"-s=bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags04.Foo, ShouldEqual, true)
		So(flags04.CommandFoo.Foo, ShouldEqual, true)
		So(flags04.CommandFoo.String, ShouldEqual, "bar")

		flags05 := struct {
			Int        int `short:"i" long:"int"`
			CommandFoo struct {
				Int        int `short:"i" long:"int"`
				CommandBar struct {
					Int int `short:"i" long:"int"`
				} `command:"bar"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-i=1",
			"foo",
			"-i=2",
			"bar",
			"-i=3",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags05.Int, ShouldEqual, 1)
		So(flags05.CommandFoo.Int, ShouldEqual, 2)
		So(flags05.CommandFoo.CommandBar.Int, ShouldEqual, 3)

		flags06 := struct {
			String     string `short:"s" long:"string"`
			CommandFoo struct {
				String string `short:"s" long:"string"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-s",
			"\"foo\"",
			"foo",
			"-s=bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags06.String, ShouldEqual, "foo")
		So(flags06.CommandFoo.String, ShouldEqual, "bar")

		flags07 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags07.Foo, ShouldEqual, true)

		flags08 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz bool `short:"b"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags08, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags08.Foo, ShouldEqual, false)
		So(flags08.CommandBar.Baz, ShouldEqual, true)

		flags09 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags09, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags09.Foo, ShouldEqual, false)
		So(flags09.CommandBar.Baz, ShouldEqual, true)

		flags10 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags10, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags10.Foo, ShouldEqual, true)
		So(flags10.CommandBar.Baz, ShouldEqual, true)

		flags11 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags11, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags11.Foo, ShouldEqual, false)
		So(flags11.CommandBar.Baz, ShouldEqual, true)
		So(flags11.CommandBar.CommandQux.Quux, ShouldEqual, false)

		flags12 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags12, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags12.Foo, ShouldEqual, true)
		So(flags12.CommandBar.Baz, ShouldEqual, true)
		So(flags12.CommandBar.CommandQux.Quux, ShouldEqual, false)

		flags13 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"qux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags13, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags13.Foo, ShouldEqual, false)
		So(flags13.CommandBar.Baz, ShouldEqual, true)
		So(flags13.CommandBar.CommandQux.Quux, ShouldEqual, false)

		flags14 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
			"qux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags14, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags14.Foo, ShouldEqual, true)
		So(flags14.CommandBar.Baz, ShouldEqual, true)
		So(flags14.CommandBar.CommandQux.Quux, ShouldEqual, false)

		flags15 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags15, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags15.Foo, ShouldEqual, false)
		So(flags15.CommandBar.Baz, ShouldEqual, true)
		So(flags15.CommandBar.CommandQux.Quux, ShouldEqual, true)

		flags16 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags16, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags16.Foo, ShouldEqual, true)
		So(flags16.CommandBar.Baz, ShouldEqual, true)
		So(flags16.CommandBar.CommandQux.Quux, ShouldEqual, true)

		flags17 := struct {
			Bool       bool `short:"b"`
			CommandFoo struct {
				Bool       bool `short:"b"`
				CommandFoo struct {
					Bool bool `short:"b"`
				} `command:"foo"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"-b",
			"foo",
			"-b",
			"foo",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags17, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags17.Bool, ShouldEqual, true)
		So(flags17.CommandFoo.Bool, ShouldEqual, true)
		So(flags17.CommandFoo.CommandFoo.Bool, ShouldEqual, true)

		flags18 := struct {
			Bool       bool `short:"b"`
			CommandFoo struct {
				Bool       bool `short:"b"`
				CommandBar struct {
					Bool       bool `short:"b"`
					CommandBaz struct {
						Bool bool `short:"b"`
					} `command:"baz"`
				} `command:"bar"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"baz",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags18, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags18.Bool, ShouldEqual, false)
		So(flags18.CommandFoo.Bool, ShouldEqual, false)
		So(flags18.CommandFoo.CommandBar.Bool, ShouldEqual, false)
	})
}

func TestFlagSet_FlagByName(t *testing.T) {
	Convey("should return a flag by the given name", t, func() {
		flags := struct {
			Foo        string `long:"foo-foo"`
			Bar        string `long:"bar.bar"`
			Baz        string `long:"baz_baz"`
			CommandFoo struct {
				Foo        string `long:"foo-foo"`
				Bar        string `long:"bar.bar"`
				Baz        string `long:"baz_baz"`
				CommandBar struct {
					Test string `long:"test"`
				} `command:"command.bar"`
			} `command:"command-foo"`
			CommandBaz struct{} `command:"command_baz"`
			Qux        string   `long:"qux"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flag := flagSet.FlagByName("Foo")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "foo-foo")
		flag = flagSet.FlagByName("Bar")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "bar.bar")
		flag = flagSet.FlagByName("Baz")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "baz_baz")
		flag = flagSet.FlagByName("CommandFoo")
		So(flag, ShouldNotBeNil)
		So(flag.Command(), ShouldEqual, "command-foo")
		flag = flagSet.FlagByName("CommandFoo.Foo")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "foo-foo")
		flag = flagSet.FlagByName("CommandFoo.Bar")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "bar.bar")
		flag = flagSet.FlagByName("CommandFoo.Baz")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "baz_baz")
		flag = flagSet.FlagByName("CommandFoo.CommandBar")
		So(flag, ShouldNotBeNil)
		So(flag.Command(), ShouldEqual, "command.bar")
		flag = flagSet.FlagByName("CommandFoo.CommandBar.Test")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "test")
		flag = flagSet.FlagByName("CommandBaz")
		So(flag, ShouldNotBeNil)
		So(flag.Command(), ShouldEqual, "command_baz")
		flag = flagSet.FlagByName("Qux")
		So(flag, ShouldNotBeNil)
		So(flag.Long(), ShouldEqual, "qux")
	})

	Convey("should return nil for the given flag name", t, func() {
		flags01 := struct {
			Test string `short:"t"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flag := flagSet.FlagByName("Foo")
		So(flag, ShouldBeNil)

		flags02 := struct {
			Test string `short:"t"`
		}{}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flag = flagSet.FlagByName("")
		So(flag, ShouldBeNil)
	})
}

func TestFlagSet_FlagArgs(t *testing.T) {
	Convey("should return flag arguments", t, func() {
		flags01 := struct {
			Foo bool `short:"f"`
			Bar bool `long:"bar"`
			Baz bool `short:"B" long:"baz"`
		}{}
		args := []string{
			"./app",
			"-f",
			"--bar",
			"--baz",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs := flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs[0], ShouldEqual, "true")
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs[0], ShouldEqual, "true")
		flagArgs = flagSet.FlagArgs("Baz")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs[0], ShouldEqual, "true")

		flags02 := struct {
			Foo []bool   `short:"f"`
			Bar []string `short:"b"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-f=false",
			"-f",
			"-b=foo",
			"-b=bar",
			"-b=baz",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 3)
		So(flagArgs, ShouldResemble, []string{"true", "false", "true"})
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 3)
		So(flagArgs, ShouldResemble, []string{"foo", "bar", "baz"})

		flags03 := struct {
			Foo bool `short:"f"`
			Bar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"true"})
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"bar"})

		flags04 := struct {
			Foo []bool `short:"f"`
			Bar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-f=false",
			"bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"true", "false"})
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"bar"})

		flags05 := struct {
			Foo bool `short:"f"`
			Bar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"bar"})

		flags06 := struct {
			Foo struct {
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"foo"})

		flags07 := struct {
			Foo struct {
				Bar bool `short:"b"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true"})

		flags08 := struct {
			Foo bool `short:"f"`
			Bar struct {
				Baz bool `short:"b"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags08, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"true"})
		flagArgs = flagSet.FlagArgs("Bar")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"bar", "-b=true"})

		flags09 := struct {
			Foo struct {
				Bar []bool `short:"b"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
			"-b=false",
			"-b=true",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags09, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 4)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true", "-b=false", "-b=true"})

		flags10 := struct {
			Foo struct {
				Bar []bool   `short:"b"`
				Baz []string `short:"z"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
			"-b=false",
			"-z=qux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags10, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 4)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true", "-b=false", "-z=qux"})

		flags11 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux int `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags11, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true"})

		flags12 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux int `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
			"baz",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags12, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true"})
		flagArgs = flagSet.FlagArgs("Foo.Baz")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"baz"})

		flags13 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux int `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-b",
			"baz",
			"-q=1",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags13, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"foo", "-b=true"})
		flagArgs = flagSet.FlagArgs("Foo.Baz")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"baz", "-q=1"})

		flags14 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux float64 `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"baz",
			"-q=0.1",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags14, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 1)
		So(flagArgs, ShouldResemble, []string{"foo"})
		flagArgs = flagSet.FlagArgs("Foo.Baz")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"baz", "-q=0.1"})

		flags15 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux string `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"baz",
			"-q=quux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags15, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo.Baz")
		So(flagArgs, ShouldBeNil)

		flags16 := struct {
			Command struct {
				Foo bool   `short:"f"`
				Bar string `long:"bar"`
			} `command:"command"`
		}{}
		args = []string{
			"./app",
			"command",
			"-f",
			"--bar=baz",
			"qux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags16, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Command")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 4)
		So(flagArgs, ShouldResemble, []string{"command", "-f=true", "--bar=baz", "qux"})

		flags17 := struct {
			Foo struct {
				Foo bool   `short:"f"`
				Bar string `long:"bar"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"-f",
			"--bar=baz",
			"qux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags17, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 4)
		So(flagArgs, ShouldResemble, []string{"foo", "-f=true", "--bar=baz", "qux"})

		flags18 := struct {
			Foo struct {
				Bar bool `short:"b"`
				Baz struct {
					Qux string `short:"q"`
				} `command:"baz"`
			} `command:"foo"`
		}{}
		args = []string{
			"./app",
			"foo",
			"--bar=baz",
			"baz",
			"--qux=quux",
		}
		flagSet, err = flagset.New(flagset.Options{Flags: &flags18, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagArgs = flagSet.FlagArgs("Foo")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"foo", "--bar=baz"})
		flagArgs = flagSet.FlagArgs("Foo.Baz")
		So(flagArgs, ShouldNotBeNil)
		So(flagArgs, ShouldHaveLength, 2)
		So(flagArgs, ShouldResemble, []string{"baz", "--qux=quux"})
	})
}

func TestFlagSet_Flags(t *testing.T) {
	Convey("should return flags", t, func() {
		flags := struct {
			Bool    bool    `short:"b" long:"bool"`
			Float64 float64 `short:"f" long:"float64"`
			Int     int     `short:"i" long:"int"`
			Int64   int64   `short:"I" long:"int64"`
			Uint    uint    `short:"u" long:"uint"`
			Uint64  uint64  `short:"U" long:"uint64"`
			String  string  `short:"s" long:"string"`
		}{}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		flagList := flagSet.Flags()
		So(flagList, ShouldNotBeNil)
		So(flagList, ShouldHaveLength, 7)
	})
}

func TestFlagSet_Errors(t *testing.T) {
	Convey("should return flag errors", t, func() {
		flags := struct {
			Bool bool `short:"b"`
		}{}
		args := []string{
			"./app",
			"-b=foo",
		}
		flagSet, err := flagset.New(flagset.Options{Flags: &flags, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("failed to parse 'foo' as bool"))
	})
}
