package main

import "fmt"
import "github.com/konstantinlevin77/urbandcli/scraper"

func main() {

	
	randomDef := scraper.GetRandomDefinition()

	fmt.Println(randomDef.Title)
	fmt.Println(randomDef.Description)

	fmt.Println()
	fmt.Println(randomDef.Example)
	fmt.Println()
	fmt.Println(randomDef.AuthorDate)
	

}