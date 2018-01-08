// A basic app
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/devfacet/gocmd"
)

func main() {
	var flags = struct {
		Help      bool `short:"h" long:"help" default:"false" description:"Display usage"`
		Version   bool `short:"v" long:"version" default:"false" description:"Display version"`
		VersionEx bool `long:"vv" default:"false" description:"Display version (extended)"`
		Echo      struct {
		} `command:"echo" description:"Print arguments"`
	}{}

	app, err := gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags:       &flags,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Version
	if flags.Version || flags.VersionEx {
		app.PrintVersion(flags.VersionEx)
		return
	}

	// Echo
	if args, ok := app.LookupFlag("Echo"); ok && args != nil {
		fmt.Printf("%s\n", strings.TrimRight(strings.TrimLeft(fmt.Sprintf("%v", args[1:]), "["), "]"))
		return
	}

	// Help
	app.PrintUsage()
	return
}
