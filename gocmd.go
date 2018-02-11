/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package gocmd is a library for building command line applications
package gocmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/devfacet/gocmd/table"

	"github.com/devfacet/gocmd/flagset"
)

// Options represents the options that can be set when creating a new command
type Options struct {
	// Name is the command name
	Name string
	// Version is the command version
	Version string
	// Description is the command description
	Description string
	// Flags hold user defined command line arguments and commands
	Flags interface{}
	// AutoHelp prints the usage content
	AutoHelp bool
	// AutoVersion prints the version content
	AutoVersion bool
	// AnyError checks all the errors and returns the first one
	AnyError bool
}

// New returns a command by the given options
func New(o Options) (*Cmd, error) {
	// Init the command
	cmd := Cmd{
		name:        o.Name,
		version:     o.Version,
		description: o.Description,
		flags:       o.Flags,
		flagSet:     &flagset.FlagSet{},
	}

	if o.Flags != nil {
		var err error
		cmd.flagSet, err = flagset.New(flagset.Options{Flags: o.Flags})
		if err != nil {
			return nil, err
		} else if o.AnyError && len(cmd.flagSet.Errors()) > 0 {
			return nil, cmd.flagSet.Errors()[0]
		}

		// Auto version
		if o.AutoVersion {
			ver := false
			if f := cmd.flagSet.FlagByArg("v", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					ver = true
				}
			}
			if f := cmd.flagSet.FlagByArg("version", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					ver = true
				}
			}
			verEx := false
			if f := cmd.flagSet.FlagByArg("vv", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					verEx = true
				}
			}

			if ver || verEx {
				cmd.PrintVersion(verEx)
				if !cmd.isTest() {
					os.Exit(0)
				}
			}
		}

		// Auto help
		if o.AutoHelp {
			help := false
			if f := cmd.flagSet.FlagByArg("h", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					help = true
				}
			}
			if f := cmd.flagSet.FlagByArg("help", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					help = true
				}
			}
			if len(os.Args) == 1 {
				help = true
			}

			if help {
				cmd.PrintUsage()
				if !cmd.isTest() {
					os.Exit(0)
				}
			}
		}
	}

	return &cmd, nil
}

// Cmd represents a command
type Cmd struct {
	name        string
	version     string
	description string
	flags       interface{}
	flagSet     *flagset.FlagSet
}

// Name returns the name of the command
func (cmd *Cmd) Name() string {
	return cmd.name
}

// Version returns the version of the command
func (cmd *Cmd) Version() string {
	return cmd.version
}

// Description returns the description of the command
func (cmd *Cmd) Description() string {
	return cmd.description
}

// LookupFlag returns the flag arguments by the given flag name
func (cmd *Cmd) LookupFlag(name string) ([]string, bool) {
	flag := cmd.flagSet.FlagByName(name)
	if flag != nil {
		return cmd.flagSet.FlagArgs(name), true
	}
	return nil, false
}

// FlagValue returns the flag value by the given flag name
func (cmd *Cmd) FlagValue(name string) interface{} {
	flag := cmd.flagSet.FlagByName(name)
	if flag != nil {
		return flag.Value()
	}
	return nil
}

// FlagArgs returns the flag arguments by the given flag name
func (cmd *Cmd) FlagArgs(name string) []string {
	if args, ok := cmd.LookupFlag(name); ok && args != nil {
		return args
	}
	return nil
}

// FlagErrors returns the list of the flag errors
func (cmd *Cmd) FlagErrors() []error {
	return cmd.flagSet.Errors()
}

// PrintVersion prints version information
func (cmd *Cmd) PrintVersion(extra bool) {
	// Init vars
	version := ""
	goVersion := runtime.Version()

	// Update Go version for tests
	if cmd.isTest() {
		goVersion = "TEST"
	}

	// Set version
	if extra == true {
		version += fmt.Sprintf("App name    : %s\n", cmd.Name())
		version += fmt.Sprintf("App version : %s\n", strings.TrimPrefix(cmd.Version(), "v"))
		version += fmt.Sprintf("Go version  : %s", goVersion)
	} else {
		version = fmt.Sprintf("%s", strings.TrimPrefix(cmd.Version(), "v"))
	}

	fmt.Println(version)
}

// PrintUsage prints usage
func (cmd *Cmd) PrintUsage() {
	fmt.Println(cmd.usageContent())
}

// usageItem represents a usage item
type usageItem struct {
	kind     string
	flagID   int
	parentID int
	left     string
	right    string
	level    int
}

// usageItems returns the list of the usage items
func (cmd *Cmd) usageItems(kind string, parentID int, level int) []*usageItem {
	// Init vars
	var result []*usageItem

	// Iterate over the flags
	for _, flag := range cmd.flagSet.Flags() {
		if flag.ParentID() != parentID {
			continue
		} else if kind != "" && flag.Kind() != kind {
			continue
		}

		level = len(flag.FieldIndex())

		if flag.Kind() == "command" {
			command := flag.Command()
			result = append(result, &usageItem{
				kind:     "command",
				flagID:   flag.ID(),
				parentID: parentID,
				left:     command,
				right:    flag.Description(),
				level:    level,
			})
			result = append(result, cmd.usageItems("", flag.ID(), level)...)
		} else if flag.Kind() == "arg" {
			arg := ""
			if flag.Short() != "" && flag.Long() != "" {
				arg = fmt.Sprintf("-%s, --%s", flag.Short(), flag.Long())
			} else if flag.Short() != "" {
				arg = fmt.Sprintf("-%s", flag.Short())
			} else if flag.Long() != "" {
				arg = fmt.Sprintf("    --%s", flag.Long())
			}
			right := flag.Description()
			def := false
			env := false
			if flag.ValueDefault() != "" && flag.ValueDefault() != "false" {
				def = true
			}
			if flag.Env() != "" {
				env = true
			}
			if def || env {
				right = fmt.Sprintf("%s (default", right)
			}
			if def {
				right = fmt.Sprintf("%s %s", right, flag.ValueDefault())
				if env {
					right = fmt.Sprintf("%s - override $%s", right, flag.Env())
				}
			} else if env {
				right = fmt.Sprintf("%s $%s", right, flag.Env())
			}
			if def || env {
				right = fmt.Sprintf("%s)", right)
			}
			result = append(result, &usageItem{
				kind:     "arg",
				flagID:   flag.ID(),
				parentID: parentID,
				left:     arg,
				right:    right,
				level:    level,
			})
		}
	}

	return result
}

// usageContent parses the flags and return the usage content
func (cmd *Cmd) usageContent() string {
	// Init vars
	hasOpt := false
	hasCmd := false
	usageItems := cmd.usageItems("", -1, 0)
	for _, v := range usageItems {
		if v.kind == "arg" {
			hasOpt = true
		} else if v.kind == "command" {
			hasCmd = true
		}
		if hasOpt && hasCmd {
			continue
		}
	}
	t := table.New(table.Options{})

	// Header and description
	usage := "Usage: " + cmd.name
	if hasOpt {
		usage += " [options...]"
	}
	if hasCmd {
		usage += " COMMAND [options...]"
	}
	usage += "\n\n"
	if cmd.description != "" {
		usage += cmd.description + "\n\n"
	}

	// Options
	if hasOpt {
		t.AddRow("Options:")
		for _, v := range usageItems {
			if v.kind == "arg" && v.parentID == -1 {
				t.AddRow(fmt.Sprintf("%s%s ", strings.Repeat("  ", v.level), v.left), v.right)
			}
		}
		t.AddRow(" ")
	}

	if hasCmd {
		t.AddRow("Commands:")
		l := len(usageItems)
		for i := 0; i < l; i++ {
			v := usageItems[i]
			if v.kind == "command" || (v.kind == "arg" && v.parentID != -1) {
				// Commands and their arguments are already sorted
				t.AddRow(fmt.Sprintf("%s%s ", strings.Repeat("  ", v.level), v.left), v.right)
			}
		}
	}

	if len(t.Data()) > 0 {
		usage += t.FormattedData()
	}

	return usage
}

func (cmd *Cmd) isTest() bool {
	if len(os.Args) > 0 {
		if strings.HasSuffix(os.Args[0], "gocmd.test") {
			return true
		}
		for _, v := range os.Args {
			if strings.HasPrefix(v, "-test.") {
				return true
			}
		}
	}
	return false
}
