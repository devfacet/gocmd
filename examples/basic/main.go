// A basic app
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help      bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version   bool `short:"v" long:"version" description:"Display version"`
		VersionEx bool `long:"vv" description:"Display version (extended)"`
		Math      struct {
			Sqrt struct {
				Number float64 `short:"n" long:"number" required:"true" description:"Number"`
			} `command:"sqrt" description:"Calculate square root"`
			Pow struct {
				Base     float64 `short:"b" long:"base" required:"true" description:"Base"`
				Exponent float64 `short:"e" long:"exponent" required:"true" description:"Exponent"`
			} `command:"pow" description:"Calculate base exponential"`
		} `command:"math" description:"Math functions"`
	}{}

	cmd, err := gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "A basic app",
		Flags:       &flags,
		AutoHelp:    true,
		AutoVersion: true,
		AnyError:    true,
	})
	if err != nil {
		log.Fatal(err)
	}

	if cmd.FlagArgs("Math") != nil {
		if cmd.FlagArgs("Math.Sqrt") != nil {
			fmt.Println(math.Sqrt(flags.Math.Sqrt.Number))
		} else if cmd.FlagArgs("Math.Pow") != nil {
			fmt.Println(math.Pow(flags.Math.Pow.Base, flags.Math.Pow.Exponent))
		} else {
			log.Fatal("invalid math command")
		}
		return
	}
}
