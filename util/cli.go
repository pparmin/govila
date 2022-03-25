package govila

import "fmt"

func Init(template string) {
	fmt.Println("Initializing project...")
	fmt.Println("Templating type: ", template)
}

func Build() {
	fmt.Println("Building project...")
}

func Help() {
	fmt.Println("Displaying help information...")
}
