/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package gocmd is a library for building command line applications
package gocmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/devfacet/gocmd/v3/flagset"
	"github.com/devfacet/gocmd/v3/table"
)

var (
	flagHandlers []*FlagHandler
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
	// Logger represents the logger that is being used for printing errors
	Logger Logger
	// ConfigType is the configuration type
	ConfigType ConfigType
	// AnyError checks all the errors and returns the first one if any
	AnyError bool
	// AutoHelp prints the usage content when the help flags are detected
	AutoHelp bool
	// AutoVersion prints the version content when the version flags are detected
	AutoVersion bool
	// ExitOnError prints the error and exits the program when there is an error
	ExitOnError bool
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
		logger:      o.Logger,
	}

	// Check the logger
	if cmd.logger == nil {
		cmd.logger = log.New(os.Stdout, "", 0)
	}

	// Check the config type
	switch o.ConfigType {
	case ConfigTypeAuto:
		o.AnyError = true
		o.AutoHelp = true
		o.AutoVersion = true
		o.ExitOnError = true
	}

	// If there is no any flag then
	if o.Flags == nil {
		return &cmd, nil
	}

	// Parse flags
	var err error
	cmd.flagSet, err = flagset.New(flagset.Options{Flags: o.Flags})
	if err != nil {
		if o.ExitOnError {
			cmd.logger.Printf("%s\n", err)
			cmd.exit(1)
		}
		return nil, err
	} else if (o.AnyError || o.ExitOnError) && len(cmd.flagSet.Errors()) > 0 {
		if o.ExitOnError {
			cmd.logger.Printf("%s\n", cmd.flagSet.Errors()[0])
			cmd.exit(1)
		}
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
		if !ver {
			if f := cmd.flagSet.FlagByArg("version", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					ver = true
				}
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
			cmd.exit(0)
		}
	}

	// Auto help
	if o.AutoHelp {
		help := false
		if len(os.Args) == 1 {
			help = true
		} else {
			if f := cmd.flagSet.FlagByArg("h", ""); f != nil {
				if v, ok := f.Value().(bool); ok && v {
					help = true
				}
			}
			if !help {
				if f := cmd.flagSet.FlagByArg("help", ""); f != nil {
					if v, ok := f.Value().(bool); ok && v {
						help = true
					}
				}
			}
		}

		if help {
			cmd.PrintUsage()
			cmd.exit(0)
		}
	}

	// Check handlers
	sort.Sort(byFlagHandlerPriority(flagHandlers))
	for _, v := range flagHandlers {
		args := cmd.FlagArgs(v.name)
		if cmd.FlagArgs(v.name) != nil {
			err := v.handler(&cmd, args)
			if err != nil {
				if v.exitOnError {
					cmd.logger.Printf("%s\n", err)
					cmd.exit(1)
				}
				return nil, err
			}
		}
	}

	return &cmd, nil
}

// ConfigType represents a configuration type
type ConfigType int

const (
	// ConfigTypeAuto is a configuration type that enables automatic functionalities
	// such as usage and version printing, exit on error, etc.
	// It sets AnyError, AutoHelp, AutoVersion, ExitOnError = true
	ConfigTypeAuto = iota + 1
)

// Cmd represents a command
type Cmd struct {
	name        string
	version     string
	description string
	flags       interface{}
	flagSet     *flagset.FlagSet
	logger      Logger
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
// Nested flags are separated by dot (i.e. Foo.Bar)
func (cmd *Cmd) LookupFlag(name string) ([]string, bool) {
	flag := cmd.flagSet.FlagByName(name)
	if flag != nil {
		return cmd.flagSet.FlagArgs(name), true
	}
	return nil, false
}

// FlagValue returns the flag value by the given flag name
// Nested flags are separated by dot (i.e. Foo.Bar)
func (cmd *Cmd) FlagValue(name string) interface{} {
	flag := cmd.flagSet.FlagByName(name)
	if flag != nil {
		return flag.Value()
	}
	return nil
}

// FlagArgs returns the flag arguments by the given flag name
// Nested flags are separated by dot (i.e. Foo.Bar)
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
		goVersion = "vTest"
	}

	// Set version
	if extra {
		version += fmt.Sprintf("App name    : %s\n", cmd.Name())
		version += fmt.Sprintf("App version : %s\n", strings.TrimPrefix(cmd.Version(), "v"))
		version += fmt.Sprintf("Go version  : %s", goVersion)
	} else {
		version = strings.TrimPrefix(cmd.Version(), "v")
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

		fil := len(flag.FieldIndex())
		if fil != 0 {
			level = len(flag.FieldIndex())
		}

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
		if strings.Contains(os.Args[0], "gocmd.test") {
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

func (cmd *Cmd) exit(code int) {
	if !cmd.isTest() {
		os.Exit(code)
	}
}

// FlagHandler represents a flag handler
type FlagHandler struct {
	// name of the flag
	// Nested flags are separated by dot (i.e. Foo.Bar)
	name string
	// priority of the flag handler
	// When there are multiple handlers for the same flag, they are sorted in descending order prior to execution.
	priority int
	// handler is the function that is called when the flag is detected
	handler func(cmd *Cmd, args []string) error
	// exitOnError prints the error and exits the program when the handler returns an error
	exitOnError bool
}

// SetPriority sets the value of the priority
func (fh *FlagHandler) SetPriority(v int) {
	fh.priority = v
}

// SetExitOnError sets the value of the exitOnError
func (fh *FlagHandler) SetExitOnError(v bool) {
	fh.exitOnError = v
}

// HandleFlag registers the flag handle for the given flag name
func HandleFlag(name string, handler func(cmd *Cmd, args []string) error) (*FlagHandler, error) {
	if name == "" {
		return nil, errors.New("invalid flag name")
	}

	// Init vars
	fh := FlagHandler{
		name:        name,
		priority:    0,
		handler:     handler,
		exitOnError: true,
	}
	flagHandlers = append(flagHandlers, &fh)

	return &fh, nil
}

// byFlagHandlerPriority implements sort.Interface for []*FlagHandler
type byFlagHandlerPriority []*FlagHandler

func (p byFlagHandlerPriority) Len() int           { return len(p) }
func (p byFlagHandlerPriority) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byFlagHandlerPriority) Less(i, j int) bool { return p[i].priority < p[j].priority }

// Logger is the interface that must be implemented by loggers
type Logger interface {
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}
