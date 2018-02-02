// A basic app
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help      bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version   bool `short:"v" long:"version" description:"Display version"`
		VersionEx bool `long:"vv" description:"Display version (extended)"`
		Echo      struct {
		} `command:"echo" description:"Print arguments"`
	}{}

	cmd, err := gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags:       &flags,
		AutoHelp:    true,
		AutoVersion: true,
	})
	if err != nil {
		log.Fatal(err)
	} else if len(cmd.FlagErrors()) > 0 {
		log.Fatal(cmd.FlagErrors()[0])
	}

	// Echo
	if args, ok := cmd.LookupFlag("Echo"); ok && args != nil {
		fmt.Printf("%s\n", strings.TrimRight(strings.TrimLeft(fmt.Sprintf("%v", args[1:]), "["), "]"))
		return
	}
}
