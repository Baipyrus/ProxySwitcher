package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a new internet proxy config",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		fmt.Print("Name: ")
		fmt.Scanln(&name)

		var command string
		fmt.Print("Command? ")
		fmt.Scanln(&command)

		fmt.Println("\nPrompting 'set' variants:")
		set := util.PromptVariants()

		fmt.Println("\nPrompting 'unset' variants:")
		unset := util.PromptVariants()

		config := util.Config{Name: name, Cmd: command, Set: set, Unset: unset}

		fmt.Println("\n\nPlease confirm the following data:")

		data, _ := json.Marshal(config)
		fmt.Printf("%s\n", string(data))

		var input string
		fmt.Print("Save this data? (Y/n) ")
		if input == "" || strings.ToLower(input) == "y" {
			util.SaveConfig(cfgPath, config)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
