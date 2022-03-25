package main

import (
	"flag"
	"fmt"
	"os"

	govila "github.com/pparmin/govila/util"
)

func main() {
	initFlags := flag.NewFlagSet("init", flag.ExitOnError)
	initType := initFlags.String("type", "", "defines the type of template to be created")

	if len(os.Args) < 2 {
		fmt.Println("No valid subcommand given: expected 'init' subcommand")
		os.Exit(1)
	}

	fmt.Println("ARGUMENT: ", os.Args[1])
	fmt.Println("FLAGS: ", os.Args[2:])

	fmt.Println()
	switch os.Args[1] {
	case "init":
		initFlags.Parse(os.Args[2:])
		fmt.Println("Subcommand init ")
		fmt.Println("	type: ", *initType)
		govila.Init(*initType)

	case "build":
		govila.Build()

	case "help":
		govila.Help()
	default:
		fmt.Println("No valid subcommand given: expected 'init' subcommand")
		os.Exit(1)
	}
}
