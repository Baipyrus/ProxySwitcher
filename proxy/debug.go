package proxy

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
)

func mapCmdsToStr(commands []*util.Command) string {
	var output []string

	for _, command := range commands {
		cmdArgs := strings.Join(command.Arguments, " ")
		cmdStr := fmt.Sprintf("%s %s", command.Name, cmdArgs)
		output = append(output, cmdStr)
	}

	return strings.Join(output, "\n")
}

func Debug() {
	proxy, _ := ReadSystemProxy()
	fmt.Println("\nSystem Proxy:")
	fmt.Printf("Enabled: %t\n", proxy.Enabled)
	fmt.Printf("Server: %s\n\n", proxy.Server)

	configs, _ := util.ReadConfigs()
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}
		fmt.Printf("Loading commands for '%s':\n", config.Name)

		setCmds := generateCommands(config.Set, configCmd, "[PROXY PLACEHOLDER]")
		fmt.Println("Set Commands:")
		fmt.Printf("%s\n", mapCmdsToStr(setCmds))

		unsetCmds := generateCommands(config.Unset, configCmd, "")
		fmt.Println("Unset Commands:")
		fmt.Printf("%s\n\n", mapCmdsToStr(unsetCmds))
	}
}
