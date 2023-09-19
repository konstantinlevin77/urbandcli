package utils

import (
	"github.com/konstantinlevin77/urbandcli/scraper"
	"github.com/fatih/color"
	"fmt"
)

func DisplayDefinition(def scraper.Definition) {

	color.New(color.FgBlue,color.Bold).Println(def.Title)
	fmt.Println()
	color.White(def.Description)
	fmt.Println()
	color.New(color.FgWhite,color.Italic).Println(def.Example)
	color.Green(def.AuthorDate)

}
