package proxy

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
)

func Set() {
	stdin, closeFunc, _ := util.ReadyCmd()

	proxy, _ := ReadSystemProxy()
	if !proxy.Enabled {
		SetSystemProxy(true)
	}

	configs, _ := util.ReadConfigs()
	for _, config := range configs {
		configCmd := config.Name
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := getVariants(config.Set, configCmd, proxy.Server)

		for _, command := range commands {
			cmdArgs := strings.Join(command.Arguments, " ")
			cmdStr := fmt.Sprintf("%s %s", command.Name, cmdArgs)

			fmt.Printf("%s\n", cmdStr)
			fmt.Fprintln(*stdin, cmdStr)
		}
	}

	closeFunc()
}
