package govila

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var (
	buf bytes.Buffer
)

func Init(path, name string) {
	fmt.Printf("Initializing project with name '%s'...\n", name)
	fmt.Printf("Provided path: '%s'\n", path)
	setup(path, name)
}

func Build() {
	fmt.Println("Building project...")
}

func Help() {
	fmt.Println("Displaying help information...")
}

func Remove() {
	fmt.Println("Removing project...")
}

func ParseMD() {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	source, err := os.ReadFile("util/testfiles/test.md")
	if err != nil {
		fmt.Print("Something went wrong while reading of file: ")
		log.Fatal(err)
	}

	dest, err := os.Create("util/testfiles/test.html")
	if err != nil {
		fmt.Print("Something went wrong while creating the file: ")
		log.Fatal(err)
	}
	// source := []byte(`# Hello World
	// * one
	// * two
	// * threeq

	// **CODE:**
	// ## Heading 3`)
	fmt.Printf("Rendering of Markdown initiated: %s\n", source)
	if err := md.Convert(source, dest); err != nil {
		fmt.Print("Something went wrong during the parsing of markdown: ")
		log.Fatal(err)
	}
	fmt.Printf("Markdown successfully parsed and written to: %s\n", buf.String())
	//err := os.WriteFile("test.html", []byte(buf))

}

func ShowDefault(dir string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	fmt.Println("DEFAULT PATH: ", wd)
}

func setup(p, n string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	if wd != p {
		fmt.Printf("Path provided is different from current directory. Changing directory to '%s'\n", p)
		err = os.Chdir(p)
		if err != nil {
			fmt.Print("Error during change of directory: ")
			log.Fatal(err)
		}
	}

	err = os.Mkdir(n, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Print("Error during creation of directory: ")
		log.Fatal(err)
	}
	wd, err = os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	fmt.Printf("Created directory '/%s' at path '%s\n'", n, wd)
	fmt.Println("Creating subdirectories...")
	os.Chdir("./" + n)
	wd, err = os.Getwd()
	if err != nil {
		fmt.Print("Error during request of working directory: ")
		log.Fatal(err)
	}
	fmt.Println("Changed directory to: ", wd)

	err = os.Mkdir("static", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Print("Error during creation of directory: ")
		log.Fatal(err)
	}
	fmt.Println("Created directory '/static' at path: ", wd)

	err = os.Mkdir("templates", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Print("Error during creation of directory: ")
		log.Fatal(err)
	}
	fmt.Println("Created directory '/templates' at path: ", wd)
}
