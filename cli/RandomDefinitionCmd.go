package cli

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/konstantinlevin77/urbandcli/scraper"
	"github.com/spf13/cobra"
)


var RandomDefinitionCmd = &cobra.Command{
	Use: "random",
	Short: "A random definition from Urban Dictionary",
	Run: func (cmd *cobra.Command, args []string)  {

		randomDef := scraper.GetRandomDefinition()
		color.New(color.FgBlue,color.Bold).Println(randomDef.Title)
		fmt.Println()
		color.White(randomDef.Description)
		fmt.Println()
		color.New(color.FgWhite,color.Italic).Println(randomDef.Example)
		color.Green(randomDef.AuthorDate)
		
	},
}