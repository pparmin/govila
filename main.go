package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	govila "github.com/pparmin/govila/util"
)

var (
	defaultPath = getDefaultPath()
	defaultName = "blog"
	initFlags   = flag.NewFlagSet("init", flag.ExitOnError)
	initPath    = initFlags.String("path", defaultPath, "specify the path for the root directory of the project")
	initName    = initFlags.String("name", defaultName, "specify the name of your project")

	buildFlags = flag.NewFlagSet("build", flag.ExitOnError)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No valid subcommand given: expected 'init' subcommand")
		os.Exit(1)
	}

	flag.CommandLine.SetOutput(os.Stdout)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	blog := govila.NewBlog(defaultName, defaultPath)

	fmt.Println("ARGUMENT: ", os.Args[1])
	fmt.Println("FLAGS: ", os.Args[2:])
	fmt.Println()
	switch os.Args[1] {
	case "init":
		initFlags.Parse(os.Args[2:])
		fmt.Println("INIT SUBCOMMAND CALLED")
		fmt.Println()

		// review pointer arguments; are they really necessary?
		blog.Setup(*initName, *initPath)

	case "build":
		/* CURRENT PROBLEM: location of blog is not retained across executions
		   --> We need a config file that tracks the blog location
		*/
		blog.Build()

	// case "help":
	// 	govila.Help()

	// case "remove":
	// 	govila.Remove()

	// case "parseMD":
	// 	govila.ParseMD()

	// case "showDefault":
	// 	govila.ShowDefault(defaultPath)

	// add help command as default output ?
	default:
		fmt.Println("No valid subcommand given: expected 'init' subcommand")
		os.Exit(1)
	}
}

func getDefaultPath() (path string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	fmt.Println("DEFAULT PATH IS: ", wd)
	return wd
}
