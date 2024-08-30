package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a new internet proxy config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Saving New Proxy Config...\n")

		// Block process until interrupted
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		fmt.Println("Blocking, press ctrl+c to continue...")
		<-done
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
