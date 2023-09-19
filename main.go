package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/konstantinlevin77/urbandcli/cli"
)

func init(){

	

}

func main() {

	var rootCmd = &cobra.Command{
		Use: "urbandcli",
		Run: func(cmd *cobra.Command, args []string) {
			color.Yellow("urbandcli is an unofficial Urban Dictionary CLI, type urbandcli --help for more.")
		},
		DisableAutoGenTag: true,
		
	}

	rootCmd.AddCommand(cli.RandomDefinitionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}