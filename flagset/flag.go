/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

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
)

// Flag represents a flag
type Flag struct {
	id           int
	name         string
	short        string
	long         string
	command      string
	description  string
	required     bool
	env          string
	delimiter    string
	valueDefault string
	valueType    string
	valueBy      string
	kind         string
	fieldIndex   []int // for reflect
	parentIndex  []int // for reflect
	parentID     int
	commandID    int
	args         []*Arg
	updatedBy    string
	err          error
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

// Kind returns the kind of the flag
func (f *Flag) Kind() string {
	return f.kind
}

// ParentID returns the parent flag id of the flag
func (f *Flag) ParentID() int {
	return f.parentID
}
