/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mtintes/cheat-indexer/actions"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var addIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory")
			return
		}
		//cheat
		//keywords: Thing, OtherThing, AnotherThing
		//description: This is a thing that does a thing
		if _, err := os.Stat(fmt.Sprintf("%s/.config/cheat-indexer/indexes.json", homedir)); os.IsNotExist(err) {
			fmt.Println("Could not find config file. Run 'cheat-indexer init' to create one")
			return
		}

		indexPathToAdd := ""
		if len(args) == 0 {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory")
				os.Exit(1)
			}
			indexPathToAdd = pwd
		} else {
			indexPathToAdd = args[0]
		}

		fmt.Println("index called")

		actions.AddIndex(indexPathToAdd, fmt.Sprintf("%s/.config/cheat-indexer/indexes.json", homedir))
	},
}

func init() {
	addCmd.AddCommand(addIndexCmd)

}
