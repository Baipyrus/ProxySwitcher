//go:build windows

package cmd

import (
	"github.com/Baipyrus/ProxySwitcher/proxy"
	"github.com/spf13/cobra"
)

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Disable the current internet proxy settings",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Unset(cfgPath)
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
