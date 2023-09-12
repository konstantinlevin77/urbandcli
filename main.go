package main

import "fmt"
import "github.com/konstantinlevin77/urbandcli/scraper"

func main() {

	fmt.Println("Main package ready!")
	
	randomDef := scraper.GetRandomDefinition()

	fmt.Println(randomDef)


}