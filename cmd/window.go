//go:build cgo

package cmd

import (
	g "github.com/AllenDang/giu"
	"github.com/spf13/cobra"
)

var output string = "N/A"

// windowCmd represents the window command
var windowCmd = &cobra.Command{
	Use:   "window",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		wnd := g.NewMasterWindow("Window", 400, 200, g.MasterWindowFlagsNotResizable)
		wnd.Run(func() {
			g.SingleWindow().Layout(
				g.Label("Test!"),
				g.Row(
					g.Button("Hello").OnClick(func() {
						output = "Hello!"
					}),
					g.Button("World").OnClick(func() {
						output = "World!"
					}),
				),
				g.Label(output),
			)
		})
	},
}

func init() {
	rootCmd.AddCommand(windowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// windowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// windowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
