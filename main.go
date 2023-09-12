package main

import "fmt"
import "github.com/konstantinlevin77/urbandcli/scraper"

func main() {

	
	testSearch := scraper.SearchDefinition("Mehmet",1)

	for _,item := range testSearch {
		fmt.Println(item.Title)
		fmt.Println(item.Description)
		fmt.Println()
		fmt.Println(item.Example)
		fmt.Println(item.AuthorDate)
	}

	

}