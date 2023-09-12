package main

import "fmt"
import "github.com/konstantinlevin77/urbandcli/scraper"

func main() {

	fmt.Println("Main package ready!")
	defs := scraper.GetRandomEntries()

	for _,def:= range defs {
		fmt.Println(def)
		fmt.Print("\n")
	}

}