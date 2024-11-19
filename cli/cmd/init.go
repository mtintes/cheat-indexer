/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mtintes/cheat-indexer/types"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
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

		if _, err := os.Stat(fmt.Sprintf("%s/.config/cheat-indexer/indexes.json", homedir)); !os.IsNotExist(err) {
			fmt.Println("Config file already exists")
			return
		}

		if _, err := os.Stat(fmt.Sprintf("%s/.config/cheat-indexer", homedir)); os.IsNotExist(err) {
			err = os.MkdirAll(fmt.Sprintf("%s/.config/cheat-indexer", homedir), 0755)

			if err != nil {
				fmt.Println("Error creating config directory")
				return
			}
		}

		configFile, err := os.Create(fmt.Sprintf("%s/.config/cheat-indexer/%s", homedir, "indexes.json"))
		if err != nil {
			fmt.Println("Error creating config file")
			return
		}
		defer configFile.Close()

		blankConfigFile := types.Config{
			Version:      "1.0",
			Repositories: []types.Repository{},
		}

		blankConfigFileJson, err := json.Marshal(blankConfigFile)
		if err != nil {
			fmt.Println("Error creating config file")
			return
		}

		_, err = configFile.WriteString(string(blankConfigFileJson))
		if err != nil {
			fmt.Println("Error creating config file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
