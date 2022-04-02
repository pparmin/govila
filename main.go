package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	govila "github.com/pparmin/govila/util"
)

var (
	defaultDir = ""
	initFlags  = flag.NewFlagSet("init", flag.ExitOnError)
	initPath   = initFlags.String("path", defaultDir, "specify the path for the root directory of the project")
	initName   = initFlags.String("name", "project", "specify the name of your project")
)

func main() {
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
		fmt.Println()
		govila.Init(*initPath, *initName)

	case "build":
		govila.Build()

	case "help":
		govila.Help()

	case "remove":
		govila.Remove()

	case "parseMD":
		govila.ParseMD()

	case "showDefault":
		govila.ShowDefault(defaultDir)

	default:
		fmt.Println("No valid subcommand given: expected 'init' subcommand")
		os.Exit(1)
	}
}

func init() {
	defaultDir = getDefaultPath()
}

func getDefaultPath() (path string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	fmt.Println("DEFAULT PATH ASSIGNED TO CURRENT PATH: ", wd)
	return wd
}
