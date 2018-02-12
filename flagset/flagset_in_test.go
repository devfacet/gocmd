/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlagSet(t *testing.T) {
	Convey("should return the correct flag set and values", t, func() {
		flags01 := struct {
			Default  bool      `short:"d" long:"default" default:"false" description:"Default argument"`
			Required bool      `short:"r" long:"required" required:"true"`
			Nonempty bool      `short:"n" long:"nonempty" nonempty:"true"`
			Global   bool      `short:"g" long:"global" global:"true"`
			Env      string    `short:"e" long:"env" env:"GOPATH"`
			Bool     bool      `short:"b" long:"bool"`
			Float64  float64   `short:"f" long:"float64"`
			Int      int       `short:"i" long:"int"`
			Int64    int64     `short:"I" long:"int64"`
			Uint     uint      `short:"u" long:"uint"`
			Uint64   uint64    `short:"U" long:"uint64"`
			String   string    `short:"s" long:"string"`
			Bools    []bool    `long:"bools"`
			Floats   []float64 `long:"floats"`
			Ints     []int     `long:"ints"`
			Int64s   []int64   `long:"int64s"`
			Uints    []uint    `long:"uints"`
			Uint64s  []uint64  `long:"uint64s"`
			Strings  []string  `long:"strings"`
			Command  struct{}  `command:"cmd"`
			NoFlag   string
		}{}
		args := []string{"./app"}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldNotBeNil)
		So(flagErrors, ShouldContain, errors.New("argument -r is required"))

		flagTests := []struct {
			id           int
			name         string
			short        string
			long         string
			command      string
			description  string
			required     bool
			nonempty     bool
			global       bool
			env          string
			delimiter    string
			valueDefault string
			valueType    string
			valueBy      string
			kind         string
			fieldIndex   []int
			parentIndex  []int
			parentID     int
			commandID    int
			args         []*Arg
			err          error
			updatedBy    []string
			value        interface{}
		}{
			{
				id:           0,
				name:         "Default",
				short:        "d",
				long:         "default",
				command:      "",
				description:  "Default argument",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "false",
				valueType:    "bool",
				valueBy:      "default",
				kind:         "arg",
				fieldIndex:   []int{0},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        false,
			},
			{
				id:           1,
				name:         "Required",
				short:        "r",
				long:         "required",
				command:      "",
				description:  "",
				required:     true,
				nonempty:     true,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "bool",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{1},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          errors.New("argument -r is required"),
				updatedBy:    nil,
				value:        false,
			},
			{
				id:           2,
				name:         "Nonempty",
				short:        "n",
				long:         "nonempty",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     true,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "bool",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{2},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        false,
			},
			{
				id:           3,
				name:         "Global",
				short:        "g",
				long:         "global",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       true,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "bool",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{3},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        false,
			},
			{
				id:           4,
				name:         "Env",
				short:        "e",
				long:         "env",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "GOPATH",
				delimiter:    "",
				valueDefault: "",
				valueType:    "string",
				valueBy:      "env",
				kind:         "arg",
				fieldIndex:   []int{4},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        os.Getenv("GOPATH"),
			},
			{
				id:           5,
				name:         "Bool",
				short:        "b",
				long:         "bool",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "bool",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{5},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        false,
			},
			{
				id:           6,
				name:         "Float64",
				short:        "f",
				long:         "float64",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "float64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{6},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        0,
			},
			{
				id:           7,
				name:         "Int",
				short:        "i",
				long:         "int",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "int",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{7},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        0,
			},
			{
				id:           8,
				name:         "Int64",
				short:        "I",
				long:         "int64",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "int64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{8},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        0,
			},
			{
				id:           9,
				name:         "Uint",
				short:        "u",
				long:         "uint",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "uint",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{9},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        0,
			},
			{
				id:           10,
				name:         "Uint64",
				short:        "U",
				long:         "uint64",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "uint64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{10},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        0,
			},
			{
				id:           11,
				name:         "String",
				short:        "s",
				long:         "string",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "string",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{11},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        "",
			},
			{
				id:           12,
				name:         "Bools",
				short:        "",
				long:         "bools",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]bool",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{12},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        nil,
			},
			{
				id:           13,
				name:         "Floats",
				short:        "",
				long:         "floats",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]float64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{13},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []float64{},
			},
			{
				id:           14,
				name:         "Ints",
				short:        "",
				long:         "ints",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]int",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{14},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []int{},
			},
			{
				id:           15,
				name:         "Int64s",
				short:        "",
				long:         "int64s",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]int64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{15},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []int64{},
			},
			{
				id:           16,
				name:         "Uints",
				short:        "",
				long:         "uints",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]uint",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{16},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []uint{},
			},
			{
				id:           17,
				name:         "Uint64s",
				short:        "",
				long:         "uint64s",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]uint64",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{17},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []uint64{},
			},
			{
				id:           18,
				name:         "Strings",
				short:        "",
				long:         "strings",
				command:      "",
				description:  "",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "[]string",
				valueBy:      "",
				kind:         "arg",
				fieldIndex:   []int{18},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    -1,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        []string{},
			},
			{
				id:           19,
				name:         "Command",
				short:        "",
				long:         "",
				description:  "",
				command:      "cmd",
				required:     false,
				nonempty:     false,
				global:       false,
				env:          "",
				delimiter:    "",
				valueDefault: "",
				valueType:    "struct",
				valueBy:      "",
				kind:         "command",
				fieldIndex:   []int{19},
				parentIndex:  nil,
				parentID:     -1,
				commandID:    0,
				args:         nil,
				err:          nil,
				updatedBy:    nil,
				value:        struct{}{},
			},
		}
		for _, v := range flagTests {
			f := flagSet.flagByID(v.id)

			So(f, ShouldNotBeNil)
			So(f.id, ShouldEqual, v.id)
			So(f.name, ShouldEqual, v.name)
			So(f.short, ShouldEqual, v.short)
			So(f.long, ShouldEqual, v.long)
			So(f.command, ShouldEqual, v.command)
			So(f.description, ShouldEqual, v.description)
			So(f.required, ShouldEqual, v.required)
			So(f.nonempty, ShouldEqual, v.nonempty)
			So(f.global, ShouldEqual, v.global)
			So(f.env, ShouldEqual, v.env)
			So(f.delimiter, ShouldEqual, v.delimiter)
			So(f.valueDefault, ShouldEqual, v.valueDefault)
			So(f.valueType, ShouldEqual, v.valueType)
			So(f.valueBy, ShouldEqual, v.valueBy)
			So(f.kind, ShouldEqual, v.kind)
			So(f.fieldIndex, ShouldResemble, v.fieldIndex)
			So(f.parentIndex, ShouldEqual, v.parentIndex)
			So(f.parentID, ShouldEqual, v.parentID)
			So(f.commandID, ShouldEqual, v.commandID)
			So(f.args, ShouldEqual, v.args)
			if v.err == nil {
				So(f.err, ShouldEqual, v.err)
			} else {
				So(f.err, ShouldBeError, v.err)
			}
			So(f.updatedBy, ShouldEqual, v.updatedBy)

			switch v.name {
			case "Default":
				So(flags01.Default, ShouldEqual, v.value)
			case "Required":
				So(flags01.Required, ShouldEqual, v.value)
			case "Nonempty":
				So(flags01.Nonempty, ShouldEqual, v.value)
			case "Global":
				So(flags01.Global, ShouldEqual, v.value)
			case "Env":
				So(flags01.Env, ShouldEqual, v.value)
			case "Bool":
				So(flags01.Bool, ShouldEqual, v.value)
			case "Float64":
				So(flags01.Float64, ShouldEqual, v.value)
			case "Int":
				So(flags01.Int, ShouldEqual, v.value)
			case "Int64":
				So(flags01.Int64, ShouldEqual, v.value)
			case "Uint":
				So(flags01.Uint, ShouldEqual, flags01.Uint)
			case "Uint64":
				So(flags01.Uint64, ShouldEqual, v.value)
			case "String":
				So(flags01.String, ShouldEqual, v.value)
			case "Bools":
				So(flags01.Bools, ShouldEqual, v.value)
			case "Floats":
				So(fmt.Sprint(flags01.Floats), ShouldEqual, fmt.Sprint(v.value))
			case "Ints":
				So(fmt.Sprint(flags01.Ints), ShouldEqual, fmt.Sprint(v.value))
			case "Int64s":
				So(fmt.Sprint(flags01.Int64s), ShouldEqual, fmt.Sprint(v.value))
			case "Uints":
				So(fmt.Sprint(flags01.Uints), ShouldEqual, fmt.Sprint(v.value))
			case "Uint64s":
				So(fmt.Sprint(flags01.Uint64s), ShouldEqual, fmt.Sprint(v.value))
			case "Strings":
				So(fmt.Sprint(flags01.Strings), ShouldEqual, fmt.Sprint(v.value))
			case "Command":
				So(fmt.Sprint(flags01.Command), ShouldEqual, fmt.Sprint(v.value))
			}
		}

		flags02 := struct {
			Foo        bool   `short:"f"`
			String     string `short:"s"`
			CommandBar struct {
				String     string `short:"s"`
				CommandQux struct {
					String   string `short:"s"`
					Settings bool   `settings:"true" allow-unknown-arg:"true"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s",
			"foo1",
			"-s=foo2",
			"bar",
			"-s=foo3",
			"qux",
			"-s=foo4",
			"quux",
		}
		flagSet, err = New(Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flags02.Foo, ShouldEqual, true)
		So(flags02.String, ShouldEqual, "foo2")
		So(flags02.CommandBar.String, ShouldEqual, "foo3")
		So(flags02.CommandBar.CommandQux.String, ShouldEqual, "foo4")
		So(flagSet.flags, ShouldHaveLength, 6)
		So(flagSet.flags[0].args, ShouldNotBeNil)
		So(flagSet.flags[0].args, ShouldHaveLength, 1)
		So(flagSet.flags[1].args, ShouldNotBeNil)
		So(flagSet.flags[1].args, ShouldHaveLength, 2)
		So(flagSet.flags[2].args, ShouldNotBeNil)
		So(flagSet.flags[2].args, ShouldHaveLength, 2)
		So(flagSet.flags[3].args, ShouldNotBeNil)
		So(flagSet.flags[3].args, ShouldHaveLength, 1)
		So(flagSet.flags[4].args, ShouldNotBeNil)
		So(flagSet.flags[4].args, ShouldHaveLength, 3)
		So(flagSet.flags[5].args, ShouldNotBeNil)
		So(flagSet.flags[5].args, ShouldHaveLength, 1)
	})
}

func TestFlagSet_settingByID(t *testing.T) {
	Convey("should return nil when the setting id is not valid", t, func() {
		flagSet, err := New(Options{Flags: &struct{}{}})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.settingByID(-1), ShouldBeNil)
	})

	Convey("should return nil when the setting id doesn't exist", t, func() {
		flags01 := struct {
			Settings bool `settings:"true"`
		}{}
		args := []string{
			"./app",
		}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.settingByID(1), ShouldBeNil)
	})
}

func TestFlagSet_commandByID(t *testing.T) {
	Convey("should return nil when the command id is not valid", t, func() {
		flagSet, err := New(Options{Flags: &struct{}{}})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.commandByID(-1), ShouldBeNil)
	})

	Convey("should return nil when the command id doesn't exist", t, func() {
		flags01 := struct {
			CommandBar struct {
			} `command:"bar"`
		}{}
		args := []string{
			"./app",
		}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.commandByID(1), ShouldBeNil)
	})
}

func TestFlagSet_flagByID(t *testing.T) {
	Convey("should return nil when the flag id is not valid", t, func() {
		flagSet, err := New(Options{Flags: &struct{}{}})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.flagByID(-1), ShouldBeNil)
	})

	Convey("should return nil when the flag id doesn't exist", t, func() {
		flagSet, err := New(Options{Flags: &struct{ NoFlag string }{}})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.flagByID(0), ShouldBeNil)

		flagSet, err = New(Options{Flags: &struct{}{}})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.flagByID(0), ShouldBeNil)
	})
}

func TestFlagSet_flagByIndex(t *testing.T) {
	Convey("should return nil when the flag index is nil", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.flagByIndex(nil), ShouldBeNil)
	})

	Convey("should return nil when the flag index doesn't exist", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.flagByIndex([]int{1}), ShouldBeNil)
	})
}

func TestFlagSet_parseSettings(t *testing.T) {
	Convey("should parse settings", t, func() {
		flagset := FlagSet{}
		So(flagset.settingsParsed, ShouldEqual, false)
		So(flagset.commandsParsed, ShouldEqual, false)
		So(flagset.argsParsed, ShouldEqual, false)
		flagset.parseSettings()
		flagset.parseSettings()
		So(flagset.settingsParsed, ShouldEqual, true)

		flags01 := struct {
			CommandBar struct {
			} `command:"bar"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
			Foo      bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args := []string{
			"./app",
		}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		flagErrors := flagSet.Errors()
		So(flagErrors, ShouldContain, errors.New("duplicate settings tag for `Foo` and `Settings` flags"))
	})
}

func TestFlagSet_parseCommands(t *testing.T) {
	Convey("should parse commands", t, func() {
		flagset := FlagSet{}
		So(flagset.commandsParsed, ShouldEqual, false)
		flagset.parseArgs()
		flagset.parseCommands()
		So(flagset.commandsParsed, ShouldEqual, true)

		flags01 := struct {
			CommandBar struct {
			} `command:"bar"`
		}{}
		args := []string{
			"./app",
		}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 1)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 0)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, -1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[0].indexTo, ShouldEqual, -1)

		flags02 := struct {
			CommandBar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
		}
		flagSet, err = New(Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 1)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 0)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 2)

		flags03 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
		}
		flagSet, err = New(Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 1)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)

		flags04 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz bool `short:"b"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
		}
		flagSet, err = New(Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 1)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)

		flags05 := struct {
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
		flagSet, err = New(Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 1)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)

		flags06 := struct {
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 2)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 2)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[1].indexTo, ShouldEqual, 3)

		flags07 := struct {
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 3)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[1].indexTo, ShouldEqual, 4)

		flags08 := struct {
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
		flagSet, err = New(Options{Flags: &flags08, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 5)

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
			"-f",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags09, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 6)

		flags11 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags11, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "bar")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 6)

		flags12 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags12, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)
		So(flagSet.commands[1].command, ShouldEqual, "bar")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 3)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[1].indexTo, ShouldEqual, 5)

		flags13 := struct {
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags13, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 0)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)
		So(flagSet.commands[1].command, ShouldEqual, "bar")
		So(flagSet.commands[1].flagID, ShouldEqual, 2)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 3)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[1].indexTo, ShouldEqual, 5)

		flags14 := struct {
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			CommandQux struct {
				Foo bool `short:"f"`
			} `command:"qux"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-q",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags14, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 3)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 0)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 1)
		So(flagSet.commands[0].indexTo, ShouldEqual, 3)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 2)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 3)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[1].indexTo, ShouldEqual, 5)
		So(flagSet.commands[2].command, ShouldEqual, "qux")
		So(flagSet.commands[2].flagID, ShouldEqual, 4)
		So(flagSet.commands[2].parentID, ShouldEqual, -1)
		So(flagSet.commands[2].argID, ShouldEqual, -1)
		So(flagSet.commands[2].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[2].indexTo, ShouldEqual, -1)

		flags15 := struct {
			Foo        string `short:"f"`
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags15, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, -1)
		So(flagSet.commands[0].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[0].indexTo, ShouldEqual, -1)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, -1)
		So(flagSet.commands[1].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[1].indexTo, ShouldEqual, -1)

		flags16 := struct {
			Foo        string `short:"f"`
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
			"qux",
			"bar",
		}
		flagSet, err = New(Options{Flags: &flags16, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 3)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, -1)
		So(flagSet.commands[1].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[1].indexTo, ShouldEqual, -1)

		flags17 := struct {
			Foo        string `short:"f"`
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
			"qux",
			"bar",
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags17, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 3)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 5)

		flags18 := struct {
			Foo        string `short:"f"`
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
			"qux",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags18, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 3)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 3)
		So(flagSet.commands[0].indexTo, ShouldEqual, 5)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 5)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 5)
		So(flagSet.commands[1].indexTo, ShouldEqual, 7)

		flags19 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			CommandBaz struct {
				Baz        bool `short:"b"`
				CommandFoo struct {
					bar bool `short:"b"`
				} `command:"foo"`
			} `command:"baz"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b=true",
			"qux",
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags19, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 4)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 6)
		So(flagSet.commands[2].command, ShouldEqual, "baz")
		So(flagSet.commands[2].flagID, ShouldEqual, 5)
		So(flagSet.commands[2].parentID, ShouldEqual, -1)
		So(flagSet.commands[2].argID, ShouldEqual, -1)
		So(flagSet.commands[2].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[2].indexTo, ShouldEqual, -1)
		So(flagSet.commands[3].command, ShouldEqual, "foo")
		So(flagSet.commands[3].flagID, ShouldEqual, 7)
		So(flagSet.commands[3].parentID, ShouldEqual, 2)
		So(flagSet.commands[3].argID, ShouldEqual, -1)
		So(flagSet.commands[3].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[3].indexTo, ShouldEqual, -1)

		flags20 := struct {
			Foo        bool   `short:"f"`
			String     string `short:"s"`
			CommandBar struct {
				String     string `short:"s"`
				CommandQux struct {
					String string `short:"s"`
				} `command:"qux"`
			} `command:"bar"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s",
			"foo1",
			"-s=foo2",
			"bar",
			"-s=foo3",
			"qux",
			"-s=foo4",
			"quux",
		}
		flagSet, err = New(Options{Flags: &flags20, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 2)
		So(flagSet.commands[0].id, ShouldEqual, 0)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 2)
		So(flagSet.commands[0].parentID, ShouldEqual, -1)
		So(flagSet.commands[0].argID, ShouldEqual, 5)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 5)
		So(flagSet.commands[0].indexTo, ShouldEqual, 7)
		So(flagSet.commands[1].id, ShouldEqual, 1)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 4)
		So(flagSet.commands[1].parentID, ShouldEqual, 0)
		So(flagSet.commands[1].argID, ShouldEqual, 7)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 7)
		So(flagSet.commands[1].indexTo, ShouldEqual, 10)

		flags99 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			CommandBaz struct {
				Baz        bool `short:"b"`
				CommandFoo struct {
					bar bool `short:"b"`
				} `command:"foo"`
			} `command:"baz"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b=true",
			"qux",
			"quux",
		}
		flagSet, err = New(Options{Flags: &flags99, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.commands, ShouldNotBeNil)
		So(flagSet.commands, ShouldHaveLength, 4)
		So(flagSet.commands[0].command, ShouldEqual, "bar")
		So(flagSet.commands[0].flagID, ShouldEqual, 1)
		So(flagSet.commands[0].argID, ShouldEqual, 2)
		So(flagSet.commands[0].indexFrom, ShouldEqual, 2)
		So(flagSet.commands[0].indexTo, ShouldEqual, 4)
		So(flagSet.commands[1].command, ShouldEqual, "qux")
		So(flagSet.commands[1].flagID, ShouldEqual, 3)
		So(flagSet.commands[1].argID, ShouldEqual, 4)
		So(flagSet.commands[1].indexFrom, ShouldEqual, 4)
		So(flagSet.commands[1].indexTo, ShouldEqual, 6)
		So(flagSet.commands[2].command, ShouldEqual, "baz")
		So(flagSet.commands[2].flagID, ShouldEqual, 5)
		So(flagSet.commands[2].argID, ShouldEqual, -1)
		So(flagSet.commands[2].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[2].indexTo, ShouldEqual, -1)
		So(flagSet.commands[3].command, ShouldEqual, "foo")
		So(flagSet.commands[3].flagID, ShouldEqual, 7)
		So(flagSet.commands[3].argID, ShouldEqual, -1)
		So(flagSet.commands[3].indexFrom, ShouldEqual, -1)
		So(flagSet.commands[3].indexTo, ShouldEqual, -1)
	})
}

func TestFlagSet_parseArgs(t *testing.T) {
	Convey("should parse arguments", t, func() {
		flagset := FlagSet{}
		So(flagset.argsParsed, ShouldEqual, false)
		So(flagset.commandsParsed, ShouldEqual, false)
		flagset.parseArgs()
		So(flagset.argsParsed, ShouldEqual, true)
		So(flagset.commandsParsed, ShouldEqual, true)

		flags01 := struct {
			Foo bool `short:"f"`
		}{}
		args := []string{
			"./app",
		}
		flagSet, err := New(Options{Flags: &flags01, Args: args})
		flagSet.parseArgs()
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 1)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)

		flags02 := struct {
			Foo bool `short:"f"`
		}{}
		args = []string{
			"./app",
			"-f",
		}
		flagSet, err = New(Options{Flags: &flags02, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 2)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)

		flags03 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
		}
		flagSet, err = New(Options{Flags: &flags03, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 3)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)

		flags04 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz bool `short:"b"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
		}
		flagSet, err = New(Options{Flags: &flags04, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 4)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)

		flags05 := struct {
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
		flagSet, err = New(Options{Flags: &flags05, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 3)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "bar")
		So(flagSet.args[1].name, ShouldEqual, "bar")
		So(flagSet.args[1].value, ShouldEqual, "")
		So(flagSet.args[1].dash, ShouldEqual, "")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "command")
		So(flagSet.args[1].flagID, ShouldEqual, 1)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].name, ShouldEqual, "b")
		So(flagSet.args[2].dash, ShouldEqual, "-")
		So(flagSet.args[2].value, ShouldEqual, "true")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "arg")
		So(flagSet.args[2].flagID, ShouldEqual, 2)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)

		flags06 := struct {
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags06, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 3)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "bar")
		So(flagSet.args[1].name, ShouldEqual, "bar")
		So(flagSet.args[1].value, ShouldEqual, "")
		So(flagSet.args[1].dash, ShouldEqual, "")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "command")
		So(flagSet.args[1].flagID, ShouldEqual, 1)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "qux")
		So(flagSet.args[2].name, ShouldEqual, "qux")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 3)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)

		flags07 := struct {
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags07, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 4)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "qux")
		So(flagSet.args[3].name, ShouldEqual, "qux")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 3)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)

		flags08 := struct {
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
		flagSet, err = New(Options{Flags: &flags08, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 5)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "-b")
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "qux")
		So(flagSet.args[4].name, ShouldEqual, "qux")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)

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
			"-f",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags09, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 6)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "-b")
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "qux")
		So(flagSet.args[4].name, ShouldEqual, "qux")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 6)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "-q")
		So(flagSet.args[5].name, ShouldEqual, "q")
		So(flagSet.args[5].value, ShouldEqual, "true")
		So(flagSet.args[5].dash, ShouldEqual, "-")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "arg")
		So(flagSet.args[5].flagID, ShouldEqual, 4)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 6)

		flags11 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags11, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 6)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "-b")
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "bar")
		So(flagSet.args[4].name, ShouldEqual, "bar")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 6)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "-q")
		So(flagSet.args[5].name, ShouldEqual, "q")
		So(flagSet.args[5].value, ShouldEqual, "true")
		So(flagSet.args[5].dash, ShouldEqual, "-")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "arg")
		So(flagSet.args[5].flagID, ShouldEqual, 4)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 6)

		flags12 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags12, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 5)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "bar")
		So(flagSet.args[1].name, ShouldEqual, "bar")
		So(flagSet.args[1].value, ShouldEqual, "")
		So(flagSet.args[1].dash, ShouldEqual, "")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "command")
		So(flagSet.args[1].flagID, ShouldEqual, 1)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "-b")
		So(flagSet.args[2].name, ShouldEqual, "b")
		So(flagSet.args[2].value, ShouldEqual, "true")
		So(flagSet.args[2].dash, ShouldEqual, "-")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "arg")
		So(flagSet.args[2].flagID, ShouldEqual, 2)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "bar")
		So(flagSet.args[3].name, ShouldEqual, "bar")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 3)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 5)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "-q")
		So(flagSet.args[4].name, ShouldEqual, "q")
		So(flagSet.args[4].value, ShouldEqual, "true")
		So(flagSet.args[4].dash, ShouldEqual, "-")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "arg")
		So(flagSet.args[4].flagID, ShouldEqual, 4)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)

		flags13 := struct {
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandBar struct {
					Quux bool `short:"q"`
				} `command:"bar"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"bar",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags13, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 5)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "bar")
		So(flagSet.args[1].name, ShouldEqual, "bar")
		So(flagSet.args[1].value, ShouldEqual, "")
		So(flagSet.args[1].dash, ShouldEqual, "")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "command")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "-b")
		So(flagSet.args[2].name, ShouldEqual, "b")
		So(flagSet.args[2].value, ShouldEqual, "true")
		So(flagSet.args[2].dash, ShouldEqual, "-")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "arg")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "bar")
		So(flagSet.args[3].name, ShouldEqual, "bar")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 5)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "-q")
		So(flagSet.args[4].name, ShouldEqual, "q")
		So(flagSet.args[4].value, ShouldEqual, "true")
		So(flagSet.args[4].dash, ShouldEqual, "-")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "arg")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)

		flags14 := struct {
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			CommandQux struct {
				Foo bool `short:"f"`
			} `command:"qux"`
		}{}
		args = []string{
			"./app",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags14, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 5)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "bar")
		So(flagSet.args[1].name, ShouldEqual, "bar")
		So(flagSet.args[1].value, ShouldEqual, "")
		So(flagSet.args[1].dash, ShouldEqual, "")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "command")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "-b")
		So(flagSet.args[2].name, ShouldEqual, "b")
		So(flagSet.args[2].value, ShouldEqual, "true")
		So(flagSet.args[2].dash, ShouldEqual, "-")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "arg")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "qux")
		So(flagSet.args[3].name, ShouldEqual, "qux")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 5)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "-q")
		So(flagSet.args[4].name, ShouldEqual, "q")
		So(flagSet.args[4].value, ShouldEqual, "true")
		So(flagSet.args[4].dash, ShouldEqual, "-")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "arg")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)

		flags15 := struct {
			Foo        string `short:"f"`
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
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags15, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 3)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "qux")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "qux")
		So(flagSet.args[2].name, ShouldEqual, "")
		So(flagSet.args[2].value, ShouldEqual, "qux")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "argval")
		So(flagSet.args[2].flagID, ShouldEqual, -1)
		So(flagSet.args[2].parentID, ShouldEqual, 1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)

		flags16 := struct {
			Foo        string `short:"f"`
			CommandBar struct {
				Baz        string `short:"b"`
				CommandQux struct {
					Quux string `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"qux",
			"bar",
		}
		flagSet, err = New(Options{Flags: &flags16, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 4)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "qux")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "qux")
		So(flagSet.args[2].name, ShouldEqual, "")
		So(flagSet.args[2].value, ShouldEqual, "qux")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "argval")
		So(flagSet.args[2].flagID, ShouldEqual, -1)
		So(flagSet.args[2].parentID, ShouldEqual, 1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "bar")
		So(flagSet.args[3].name, ShouldEqual, "bar")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 1)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)

		flags17 := struct {
			Foo        string `short:"f"`
			CommandBar struct {
				Baz        string `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
		}{}
		args = []string{
			"./app",
			"-f",
			"qux",
			"bar",
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags17, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 5)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "qux")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "qux")
		So(flagSet.args[2].name, ShouldEqual, "")
		So(flagSet.args[2].value, ShouldEqual, "qux")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "argval")
		So(flagSet.args[2].flagID, ShouldEqual, -1)
		So(flagSet.args[2].parentID, ShouldEqual, 1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "bar")
		So(flagSet.args[3].name, ShouldEqual, "bar")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 1)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "qux")
		So(flagSet.args[4].name, ShouldEqual, "qux")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)

		flags18 := struct {
			Foo        string `short:"f"`
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
			"qux",
			"bar",
			"-b",
			"qux",
			"-q",
		}
		flagSet, err = New(Options{Flags: &flags18, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 7)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "qux")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 3)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "qux")
		So(flagSet.args[2].name, ShouldEqual, "")
		So(flagSet.args[2].value, ShouldEqual, "qux")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "argval")
		So(flagSet.args[2].flagID, ShouldEqual, -1)
		So(flagSet.args[2].parentID, ShouldEqual, 1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 3)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "bar")
		So(flagSet.args[3].name, ShouldEqual, "bar")
		So(flagSet.args[3].value, ShouldEqual, "")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "command")
		So(flagSet.args[3].flagID, ShouldEqual, 1)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 5)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "-b")
		So(flagSet.args[4].name, ShouldEqual, "b")
		So(flagSet.args[4].value, ShouldEqual, "true")
		So(flagSet.args[4].dash, ShouldEqual, "-")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "arg")
		So(flagSet.args[4].flagID, ShouldEqual, 2)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "qux")
		So(flagSet.args[5].name, ShouldEqual, "qux")
		So(flagSet.args[5].value, ShouldEqual, "")
		So(flagSet.args[5].dash, ShouldEqual, "")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "command")
		So(flagSet.args[5].flagID, ShouldEqual, 3)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 7)
		So(flagSet.args[6].id, ShouldEqual, 6)
		So(flagSet.args[6].arg, ShouldEqual, "-q")
		So(flagSet.args[6].name, ShouldEqual, "q")
		So(flagSet.args[6].value, ShouldEqual, "true")
		So(flagSet.args[6].dash, ShouldEqual, "-")
		So(flagSet.args[6].hasEq, ShouldEqual, false)
		So(flagSet.args[6].unset, ShouldEqual, false)
		So(flagSet.args[6].kind, ShouldEqual, "arg")
		So(flagSet.args[6].flagID, ShouldEqual, 4)
		So(flagSet.args[6].parentID, ShouldEqual, -1)
		So(flagSet.args[6].indexFrom, ShouldEqual, 6)
		So(flagSet.args[6].indexTo, ShouldEqual, 7)

		flags19 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b=true",
			"qux",
			"qux",
		}
		flagSet, err = New(Options{Flags: &flags19, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 6)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "-b=true")
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].hasEq, ShouldEqual, true)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "qux")
		So(flagSet.args[4].name, ShouldEqual, "qux")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 6)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "qux")
		So(flagSet.args[5].name, ShouldEqual, "qux")
		So(flagSet.args[5].value, ShouldEqual, "")
		So(flagSet.args[5].dash, ShouldEqual, "")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "arg")
		So(flagSet.args[5].flagID, ShouldEqual, -1)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 6)

		flags20 := struct {
			Foo        bool   `short:"f"`
			String     string `short:"s"`
			CommandBar struct {
				String     string `short:"s"`
				CommandQux struct {
					String string `short:"s"`
				} `command:"qux"`
			} `command:"bar"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"-s",
			"foo1",
			"-s=foo2",
			"bar",
			"-s=foo3",
			"qux",
			"-s=foo4",
			"quux",
		}
		flagSet, err = New(Options{Flags: &flags20, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 10)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].commandID, ShouldEqual, -1)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].id, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "-s")
		So(flagSet.args[2].name, ShouldEqual, "s")
		So(flagSet.args[2].value, ShouldEqual, "foo1")
		So(flagSet.args[2].dash, ShouldEqual, "-")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "arg")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].commandID, ShouldEqual, -1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "foo1")
		So(flagSet.args[3].name, ShouldEqual, "")
		So(flagSet.args[3].value, ShouldEqual, "foo1")
		So(flagSet.args[3].dash, ShouldEqual, "")
		So(flagSet.args[3].hasEq, ShouldEqual, false)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "argval")
		So(flagSet.args[3].flagID, ShouldEqual, -1)
		So(flagSet.args[3].commandID, ShouldEqual, -1)
		So(flagSet.args[3].parentID, ShouldEqual, 2)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "-s=foo2")
		So(flagSet.args[4].name, ShouldEqual, "s")
		So(flagSet.args[4].value, ShouldEqual, "foo2")
		So(flagSet.args[4].dash, ShouldEqual, "-")
		So(flagSet.args[4].hasEq, ShouldEqual, true)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "arg")
		So(flagSet.args[4].flagID, ShouldEqual, 1)
		So(flagSet.args[4].commandID, ShouldEqual, -1)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 5)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "bar")
		So(flagSet.args[5].name, ShouldEqual, "bar")
		So(flagSet.args[5].value, ShouldEqual, "")
		So(flagSet.args[5].dash, ShouldEqual, "")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "command")
		So(flagSet.args[5].flagID, ShouldEqual, 2)
		So(flagSet.args[5].commandID, ShouldEqual, 0)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 7)
		So(flagSet.args[6].id, ShouldEqual, 6)
		So(flagSet.args[6].arg, ShouldEqual, "-s=foo3")
		So(flagSet.args[6].name, ShouldEqual, "s")
		So(flagSet.args[6].value, ShouldEqual, "foo3")
		So(flagSet.args[6].dash, ShouldEqual, "-")
		So(flagSet.args[6].hasEq, ShouldEqual, true)
		So(flagSet.args[6].unset, ShouldEqual, false)
		So(flagSet.args[6].kind, ShouldEqual, "arg")
		So(flagSet.args[6].flagID, ShouldEqual, 3)
		So(flagSet.args[6].commandID, ShouldEqual, 0)
		So(flagSet.args[6].parentID, ShouldEqual, -1)
		So(flagSet.args[6].indexFrom, ShouldEqual, 6)
		So(flagSet.args[6].indexTo, ShouldEqual, 7)
		So(flagSet.args[7].id, ShouldEqual, 7)
		So(flagSet.args[7].arg, ShouldEqual, "qux")
		So(flagSet.args[7].name, ShouldEqual, "qux")
		So(flagSet.args[7].value, ShouldEqual, "")
		So(flagSet.args[7].dash, ShouldEqual, "")
		So(flagSet.args[7].hasEq, ShouldEqual, false)
		So(flagSet.args[7].unset, ShouldEqual, false)
		So(flagSet.args[7].kind, ShouldEqual, "command")
		So(flagSet.args[7].flagID, ShouldEqual, 4)
		So(flagSet.args[7].commandID, ShouldEqual, 1)
		So(flagSet.args[7].parentID, ShouldEqual, -1)
		So(flagSet.args[7].indexFrom, ShouldEqual, 7)
		So(flagSet.args[7].indexTo, ShouldEqual, 10)
		So(flagSet.args[8].id, ShouldEqual, 8)
		So(flagSet.args[8].arg, ShouldEqual, "-s=foo4")
		So(flagSet.args[8].name, ShouldEqual, "s")
		So(flagSet.args[8].value, ShouldEqual, "foo4")
		So(flagSet.args[8].dash, ShouldEqual, "-")
		So(flagSet.args[8].hasEq, ShouldEqual, true)
		So(flagSet.args[8].unset, ShouldEqual, false)
		So(flagSet.args[8].kind, ShouldEqual, "arg")
		So(flagSet.args[8].flagID, ShouldEqual, 5)
		So(flagSet.args[8].commandID, ShouldEqual, 1)
		So(flagSet.args[8].parentID, ShouldEqual, -1)
		So(flagSet.args[8].indexFrom, ShouldEqual, 8)
		So(flagSet.args[8].indexTo, ShouldEqual, 9)
		So(flagSet.args[9].id, ShouldEqual, 9)
		So(flagSet.args[9].arg, ShouldEqual, "quux")
		So(flagSet.args[9].name, ShouldEqual, "quux")
		So(flagSet.args[9].value, ShouldEqual, "")
		So(flagSet.args[9].dash, ShouldEqual, "")
		So(flagSet.args[9].hasEq, ShouldEqual, false)
		So(flagSet.args[9].unset, ShouldEqual, false)
		So(flagSet.args[9].kind, ShouldEqual, "arg")
		So(flagSet.args[9].flagID, ShouldEqual, -1)
		So(flagSet.args[9].commandID, ShouldEqual, 1)
		So(flagSet.args[9].parentID, ShouldEqual, -1)
		So(flagSet.args[9].indexFrom, ShouldEqual, 9)
		So(flagSet.args[9].indexTo, ShouldEqual, 10)

		flags99 := struct {
			Foo        bool `short:"f"`
			CommandBar struct {
				Baz        bool `short:"b"`
				CommandQux struct {
					Quux bool `short:"q"`
				} `command:"qux"`
			} `command:"bar"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		}{}
		args = []string{
			"./app",
			"-f",
			"bar",
			"-b=true",
			"qux",
			"quux",
		}
		flagSet, err = New(Options{Flags: &flags99, Args: args})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.Errors(), ShouldBeNil)
		So(flagSet.args, ShouldNotBeNil)
		So(flagSet.args, ShouldHaveLength, 6)
		So(flagSet.args[0].id, ShouldEqual, 0)
		So(flagSet.args[0].arg, ShouldEqual, "./app")
		So(flagSet.args[0].name, ShouldEqual, "./app")
		So(flagSet.args[0].value, ShouldEqual, "")
		So(flagSet.args[0].dash, ShouldEqual, "")
		So(flagSet.args[0].hasEq, ShouldEqual, false)
		So(flagSet.args[0].unnamed, ShouldEqual, true)
		So(flagSet.args[0].unset, ShouldEqual, false)
		So(flagSet.args[0].kind, ShouldEqual, "arg")
		So(flagSet.args[0].flagID, ShouldEqual, -1)
		So(flagSet.args[0].commandID, ShouldEqual, -1)
		So(flagSet.args[0].parentID, ShouldEqual, -1)
		So(flagSet.args[0].valueID, ShouldEqual, -1)
		So(flagSet.args[0].indexFrom, ShouldEqual, 0)
		So(flagSet.args[0].indexTo, ShouldEqual, 1)
		So(flagSet.args[0].updatedBy, ShouldEqual, nil)
		So(flagSet.args[0].err, ShouldEqual, nil)
		So(flagSet.args[1].id, ShouldEqual, 1)
		So(flagSet.args[1].arg, ShouldEqual, "-f")
		So(flagSet.args[1].name, ShouldEqual, "f")
		So(flagSet.args[1].value, ShouldEqual, "true")
		So(flagSet.args[1].dash, ShouldEqual, "-")
		So(flagSet.args[1].hasEq, ShouldEqual, false)
		So(flagSet.args[1].unset, ShouldEqual, false)
		So(flagSet.args[1].kind, ShouldEqual, "arg")
		So(flagSet.args[1].flagID, ShouldEqual, 0)
		So(flagSet.args[1].parentID, ShouldEqual, -1)
		So(flagSet.args[1].indexFrom, ShouldEqual, 1)
		So(flagSet.args[1].indexTo, ShouldEqual, 2)
		So(flagSet.args[2].arg, ShouldEqual, "bar")
		So(flagSet.args[2].name, ShouldEqual, "bar")
		So(flagSet.args[2].value, ShouldEqual, "")
		So(flagSet.args[2].dash, ShouldEqual, "")
		So(flagSet.args[2].hasEq, ShouldEqual, false)
		So(flagSet.args[2].unset, ShouldEqual, false)
		So(flagSet.args[2].kind, ShouldEqual, "command")
		So(flagSet.args[2].flagID, ShouldEqual, 1)
		So(flagSet.args[2].parentID, ShouldEqual, -1)
		So(flagSet.args[2].indexFrom, ShouldEqual, 2)
		So(flagSet.args[2].indexTo, ShouldEqual, 4)
		So(flagSet.args[3].id, ShouldEqual, 3)
		So(flagSet.args[3].arg, ShouldEqual, "-b=true")
		So(flagSet.args[3].name, ShouldEqual, "b")
		So(flagSet.args[3].value, ShouldEqual, "true")
		So(flagSet.args[3].dash, ShouldEqual, "-")
		So(flagSet.args[3].hasEq, ShouldEqual, true)
		So(flagSet.args[3].unset, ShouldEqual, false)
		So(flagSet.args[3].kind, ShouldEqual, "arg")
		So(flagSet.args[3].flagID, ShouldEqual, 2)
		So(flagSet.args[3].parentID, ShouldEqual, -1)
		So(flagSet.args[3].indexFrom, ShouldEqual, 3)
		So(flagSet.args[3].indexTo, ShouldEqual, 4)
		So(flagSet.args[4].id, ShouldEqual, 4)
		So(flagSet.args[4].arg, ShouldEqual, "qux")
		So(flagSet.args[4].name, ShouldEqual, "qux")
		So(flagSet.args[4].value, ShouldEqual, "")
		So(flagSet.args[4].dash, ShouldEqual, "")
		So(flagSet.args[4].hasEq, ShouldEqual, false)
		So(flagSet.args[4].unset, ShouldEqual, false)
		So(flagSet.args[4].kind, ShouldEqual, "command")
		So(flagSet.args[4].flagID, ShouldEqual, 3)
		So(flagSet.args[4].parentID, ShouldEqual, -1)
		So(flagSet.args[4].indexFrom, ShouldEqual, 4)
		So(flagSet.args[4].indexTo, ShouldEqual, 6)
		So(flagSet.args[5].id, ShouldEqual, 5)
		So(flagSet.args[5].arg, ShouldEqual, "quux")
		So(flagSet.args[5].name, ShouldEqual, "quux")
		So(flagSet.args[5].value, ShouldEqual, "")
		So(flagSet.args[5].dash, ShouldEqual, "")
		So(flagSet.args[5].hasEq, ShouldEqual, false)
		So(flagSet.args[5].unset, ShouldEqual, false)
		So(flagSet.args[5].kind, ShouldEqual, "arg")
		So(flagSet.args[5].flagID, ShouldEqual, -1)
		So(flagSet.args[5].parentID, ShouldEqual, -1)
		So(flagSet.args[5].indexFrom, ShouldEqual, 5)
		So(flagSet.args[5].indexTo, ShouldEqual, 6)
	})
}

func TestFlagSet_setFlag(t *testing.T) {
	Convey("should return error when the flag id is not valid", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.setFlag(-1, ""), ShouldBeError, errors.New("flag id is required"))
	})

	Convey("should return error when the flag id doesn't exist", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.setFlag(0, ""), ShouldBeError, errors.New("no flag for id 0"))
	})

	Convey("should return error when the flag type is not supported", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		flag := Flag{fieldIndex: []int{0}}
		flagSet.flagsRaw = &struct{ Foo interface{} }{}
		flagSet.flags = append(flagSet.flags, &flag)
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.setFlag(0, ""), ShouldBeError, fmt.Errorf("invalid type . Supported types: %s", supportedFlagValueTypes))
	})

	Convey("should return error when the flag can't be set", t, func() {
		flags := struct{}{}
		flagSet, err := New(Options{Flags: &flags})
		flag := Flag{fieldIndex: []int{0}}
		flagSet.flagsRaw = &struct {
			bar struct{}
		}{}
		flagSet.flags = append(flagSet.flags, &flag)
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.setFlag(0, ""), ShouldBeError, fmt.Errorf("flag  can't be set"))
	})
}

func TestFlagSet_unsetFlag(t *testing.T) {
	Convey("should fail to unset flag", t, func() {
		flags01 := struct {
			Foo bool `short:"f"`
		}{}
		flagSet, err := New(Options{Flags: &flags01})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.unsetFlag(-1), ShouldBeError, errors.New("flag id is required"))
		So(flagSet.unsetFlag(99), ShouldBeError, errors.New("no flag for id 99"))

		flags02 := struct {
			Foo struct{} `command:"foo"`
		}{}
		flagSet, err = New(Options{Flags: &flags02})
		So(err, ShouldBeNil)
		So(flagSet, ShouldNotBeNil)
		So(flagSet.unsetFlag(0), ShouldBeError, fmt.Errorf("invalid type struct. Supported types: %s", supportedFlagValueTypes))
	})
}

func Test_structToFlags(t *testing.T) {
}

func Test_structFieldToFlag(t *testing.T) {
}

func Test_typeToStructField(t *testing.T) {
	Convey("should return nil when the value is nil", t, func() {
		So(typeToStructField(nil, nil), ShouldBeNil)
	})
}

func Test_checkFlags(t *testing.T) {
}
