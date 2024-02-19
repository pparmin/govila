package govila

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

type Config struct {
	Name     string
	Location string
}

type Blog struct {
	Name     string
	Location string
}

func NewBlog(name, path string) *Blog {
	b := Blog{
		Name:     name,
		Location: path,
	}
	return &b
}

func (b *Blog) Setup(name, path string) {
	// check if name & path changed from default values
	if b.Name != name {
		b.Name = name
	}
	if b.Location != path {
		b.Location = path
	}

	rootDir := filepath.Join(b.Location, b.Name)
	fmt.Printf("directory of project: %s\n", rootDir)
	b.Location = rootDir
	fmt.Printf("Blog location set to %s\n", b.Location)

	fmt.Printf("Creating root directory for project %s in the following directory: %s\n", b.Name, b.Location)
	err := os.MkdirAll(b.Location, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Print("Error during creation of directory: ")
		log.Fatal(err)
	} else if os.IsExist(err) {
		fmt.Print("A blog already exists at this location. Please select a different location")
		log.Fatal(err)
	}

	// possibly rename to "content" later
	pagesDir := filepath.Join(rootDir, "pages")
	err = os.MkdirAll(pagesDir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of directory: %s\n", pagesDir)
		log.Fatal(err)
	}
	indexFile := filepath.Join(pagesDir, "index.md")
	err = os.WriteFile(indexFile, []byte(INDEX_MD), 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of file: %s\n", indexFile)
		log.Fatal(err)
	}

	staticDir := filepath.Join(rootDir, "static")
	err = os.MkdirAll(staticDir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of directory: %s\n", staticDir)
		log.Fatal(err)
	}
	mainJS := filepath.Join(staticDir, "main.js")
	err = os.WriteFile(mainJS, []byte(MAIN_JS), 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of file: %s\n", mainJS)
		log.Fatal(err)
	}

	templatesDir := filepath.Join(rootDir, "templates")
	err = os.MkdirAll(templatesDir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of directory: %s\n", templatesDir)
		log.Fatal(err)
	}

	defaultHTML := filepath.Join(templatesDir, "default.html")
	err = os.WriteFile(defaultHTML, []byte(DEFAULT_HTML), 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of file: %s\n", defaultHTML)
		log.Fatal(err)
	}

	b.writeConfig()
}

func (b *Blog) Build() {
	var buf bytes.Buffer

	blogLocation := readConfig()
	pagesDir := filepath.Join(blogLocation, "pages")
	publicDir := filepath.Join(blogLocation, "public")
	err := os.MkdirAll(publicDir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of directory: %s\n", publicDir)
		log.Fatal(err)
	}

	files, err := os.ReadDir(pagesDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Printing from directory: ", pagesDir)
	for _, file := range files {
		fullPath := filepath.Join(pagesDir, file.Name())
		fmt.Println("Full path for file: ", fullPath)
		// we just need to process markdown files for now
		if strings.HasSuffix(file.Name(), "md") {

			fileSuffix, _ := strings.CutSuffix(file.Name(), ".md")
			content, err := os.ReadFile(fullPath)
			if err != nil {
				log.Fatal("Error while reading markdown file: ", err)
			}
			if err := goldmark.Convert([]byte(content), &buf); err != nil {
				log.Fatal("Error while parsing markdown file: ", err)
			}

			processedFileName := strings.Join([]string{fileSuffix, "html"}, ".")
			processedFile := filepath.Join(publicDir, processedFileName)
			err = os.WriteFile(processedFile, buf.Bytes(), 0755)
			if err != nil && !os.IsExist(err) {
				fmt.Printf("Error during creation of file: %s\n", processedFileName)
				log.Fatal(err)
			}
		} else {
			continue
		}
	}

	if err := goldmark.Convert([]byte(TESTMD), &buf); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func (b *Blog) writeConfig() {
	blogData, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(blogData))

	homeDirectory := os.Getenv("HOME")
	configDirectory := filepath.Join(homeDirectory, ".config", "govila")
	_, err = os.ReadDir(configDirectory)
	if os.IsExist(err) {
		fmt.Println("directory already exists!")
	}

	err = os.MkdirAll(configDirectory, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of directory: %s\n", configDirectory)
		log.Fatal(err)
	}
	fmt.Println("config directory successfully created at ", configDirectory)
	configFile := filepath.Join(configDirectory, "config.json")

	err = os.WriteFile(configFile, []byte(blogData), 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error during creation of file: %s\n", configFile)
		log.Fatal(err)
	}
	fmt.Println("Written file to ", configFile)
}

func readConfig() string {
	homeDirectory := os.Getenv("HOME")
	configFile := filepath.Join(homeDirectory, ".config", "govila", "config.json")
	config, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error during reading of config file: ", err)
	}

	var target map[string]string
	if err = json.Unmarshal(config, &target); err != nil {
		log.Fatal("Error during unmarshaling of config file: ", err)
	}

	blogLocation := target["Location"]
	return blogLocation
}
