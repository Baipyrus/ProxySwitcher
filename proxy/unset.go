package proxy

import (
	"github.com/Baipyrus/ProxySwitcher/util"
)

func Unset() {
	stdin, closeFunc, _ := util.ReadyCmd()

	// Unset system proxy, if not already
	proxy, _ := ReadSystemProxy()
	if proxy.Enabled {
		SetSystemProxy(false)
	}

	configs, _ := util.ReadConfigs()
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(config.Unset, configCmd, "")
		util.ExecCmds(commands, stdin)
	}

	closeFunc()
}
