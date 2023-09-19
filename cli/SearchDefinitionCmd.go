package cli

import (
	"github.com/konstantinlevin77/urbandcli/cli/utils"
	"github.com/konstantinlevin77/urbandcli/scraper"
	"github.com/spf13/cobra"
	"github.com/fatih/color"	
)

var NumSearchDefinitions int

var SearchDefinitionCmd = &cobra.Command{
	Use:"search",
	Short:"Search for something in Urban Dictionary",
	Args:cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args[] string){
		
		if NumSearchDefinitions < 1 {
			color.Red("You can display one definition at least.")
			return
		}

		if NumSearchDefinitions > 7 {
			color.Red("You can display 7 definitions at most.")
			return
		}

		defSlice := scraper.SearchDefinition(args[0],NumSearchDefinitions)

		if len(defSlice) == 0{
			color.Red("Search term not found :(")
		} else {
			for _,def := range defSlice {
				utils.DisplayDefinition(def)
			}
		}
	},

}