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
}

// New returns a command by the given options
func New(o Options) (*Cmd, error) {
	// Init the command
	cmd := Cmd{
		name:        o.Name,
		version:     o.Version,
		description: o.Description,
		flags:       o.Flags,
	}

	if o.Flags != nil {
		var err error
		cmd.flagSet, err = flagset.New(flagset.Options{Flags: o.Flags})
		if err != nil {
			return nil, err
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

// LookupFlag returns the values of the flag by the given name
func (cmd *Cmd) LookupFlag(name string) ([]string, bool) {
	flag := cmd.flagSet.FlagByName(name)
	if flag != nil {
		return cmd.flagSet.FlagArgs(name), true
	}
	return nil, false
}

// PrintVersion prints version information
func (cmd *Cmd) PrintVersion(extra bool) {
	// Init vars
	version := ""
	goVersion := runtime.Version()

	// Update Go version for tests
	for _, v := range os.Args {
		if strings.Index(v, "-test") == 0 {
			goVersion = "TEST"
			break
		}
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
	if cmd.flagSet == nil {
		return nil
	}

	// Init vars
	var result []*usageItem

	// Iterate over the flags
	for _, flag := range cmd.flagSet.Flags() {
		if flag.ParentID() != parentID {
			continue
		} else if kind != "" && flag.Kind() != kind {
			continue
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
			level++
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
			if flag.ValueDefault() != "" && flag.ValueDefault() != "false" {
				right = fmt.Sprintf("%s (default %s)", right, flag.ValueDefault())
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
	aligns := map[int]int{}
	usageItems := cmd.usageItems("", -1, 0)
	for _, v := range usageItems {
		if v.kind == "arg" {
			hasOpt = true
		} else if v.kind == "command" {
			hasCmd = true
		}
		if a, ok := aligns[v.parentID]; !ok {
			aligns[v.parentID] = len(v.left)
		} else if len(v.left) > a {
			aligns[v.parentID] = len(v.left)
		}
	}

	// Header and description
	usage := "\nUsage: " + cmd.name
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
		usage += "Options:\n"
		for _, v := range usageItems {
			align := fmt.Sprintf("%d", aligns[v.parentID])
			if v.kind == "arg" && v.parentID == -1 {
				usage += fmt.Sprintf("  %-"+align+"s\t\t%s\n", v.left, v.right)
			}
		}
		usage += "\n"
	}

	// Commands
	if hasCmd {
		usage += "Commands:\n"
		l := len(usageItems)
		for i := 0; i < l; i++ {
			v := usageItems[i]
			space := strings.Repeat("  ", v.level*2)
			align := fmt.Sprintf("%d", aligns[v.parentID])
			if v.kind == "command" {
				usage += fmt.Sprintf("  %s%-"+align+"s\t%s\n", space, v.left, v.right)
			} else if v.kind == "arg" && v.parentID != -1 {
				usage += fmt.Sprintf("  %s%-"+align+"s\t%s\n", space, v.left, v.right)
			}
		}
	}

	return usage
}
