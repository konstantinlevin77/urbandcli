package cli

import (
	"github.com/konstantinlevin77/urbandcli/cli/utils"
	"github.com/konstantinlevin77/urbandcli/scraper"
	"github.com/spf13/cobra"
	
)


var RandomDefinitionCmd = &cobra.Command{
	Use: "random",
	Short: "A random definition from Urban Dictionary",
	Run: func (cmd *cobra.Command, args []string)  {

		randomDef := scraper.GetRandomDefinition()
		utils.DisplayDefinition(randomDef)
		
		
	},
}