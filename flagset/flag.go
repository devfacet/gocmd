/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

import (
	"fmt"
)

var (
	supportedFlagTypes = []string{
		"bool",
		"float64",
		"int",
		"int64",
		"uint",
		"uint64",
		"string",
		"[]bool",
		"[]float64",
		"[]int",
		"[]int64",
		"[]uint",
		"[]uint64",
		"[]string",
		"struct",
	}
	supportedFlagValueTypes = []string{
		"bool",
		"float64",
		"int",
		"int64",
		"uint",
		"uint64",
		"string",
		"[]bool",
		"[]float64",
		"[]int",
		"[]int64",
		"[]uint",
		"[]uint64",
		"[]string",
	}
)

// Flag represents a flag
type Flag struct {
	id              int
	name            string
	short           string
	long            string
	command         string
	description     string
	required        bool // flag must be present
	nonempty        bool // if the flag is present then it must have a value
	allowUnknownArg bool // allow unknown arguments to be present
	global          bool
	delimiter       string
	env             string
	valueDefault    string
	valueType       string
	valueBy         string
	value           interface{}
	kind            string
	fieldIndex      []int // for reflect
	parentIndex     []int // for reflect
	parentID        int
	commandID       int
	args            []*Arg
	err             error
	updatedBy       []string // for debug
}

// ID returns the id of the flag
func (f *Flag) ID() int {
	return f.id
}

// Name returns the name of the flag
func (f *Flag) Name() string {
	return f.name
}

// Short returns the short argument of the flag
func (f *Flag) Short() string {
	return f.short
}

// Long returns the long argument of the flag
func (f *Flag) Long() string {
	return f.long
}

// FormattedArg returns the formatted argument of the flag
func (f *Flag) FormattedArg() string {
	arg := ""
	if f.short != "" {
		arg = fmt.Sprintf("-%s", f.short)
	}
	if f.long != "" {
		if arg != "" {
			arg = fmt.Sprintf("%s (--%s)", arg, f.long)
		} else {
			arg = fmt.Sprintf("%s%s", "--", f.long)
		}
	}
	return arg
}

// Command returns the command of the flag
func (f *Flag) Command() string {
	return f.command
}

// Description returns the description of the flag
func (f *Flag) Description() string {
	return f.description
}

// Required returns whether the flag is required or not
func (f *Flag) Required() bool {
	return f.required
}

// Nonempty returns whether the flag requires a non-empty argument value or not
func (f *Flag) Nonempty() bool {
	return f.nonempty
}

// Global returns whether the flag is global or not
func (f *Flag) Global() bool {
	return f.global
}

// Env returns the environment variable name of the flag
func (f *Flag) Env() string {
	return f.env
}

// Delimiter returns the delimiter value of the flag
func (f *Flag) Delimiter() string {
	return f.delimiter
}

// ValueDefault returns the default value of the flag
func (f *Flag) ValueDefault() string {
	return f.valueDefault
}

// ValueType returns the value type of the flag
func (f *Flag) ValueType() string {
	return f.valueType
}

// ValueBy returns the value by of the flag
func (f *Flag) ValueBy() string {
	return f.valueBy
}

// Value returns the value of the flag
func (f *Flag) Value() interface{} {
	return f.value
}

// Kind returns the kind of the flag
func (f *Flag) Kind() string {
	return f.kind
}

// FieldIndex returns the parent flag id of the flag
func (f *Flag) FieldIndex() []int {
	return f.fieldIndex
}

// ParentIndex returns the parent flag id of the flag
func (f *Flag) ParentIndex() []int {
	return f.parentIndex
}

// ParentID returns the parent flag id of the flag
func (f *Flag) ParentID() int {
	return f.parentID
}

// CommandID returns the command flag id of the flag
func (f *Flag) CommandID() int {
	return f.commandID
}

// Err returns the error of the flag
func (f *Flag) Err() error {
	return f.err
}
