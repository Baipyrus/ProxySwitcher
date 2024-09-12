package proxy

import (
	"github.com/Baipyrus/ProxySwitcher/util"
)

func Set() {
	stdin, closeFunc, _ := util.ReadyCmd()

	proxy, _ := ReadSystemProxy()
	// Set system proxy, if not already
	if !proxy.Enabled {
		SetSystemProxy(true)
	}

	configs, _ := util.ReadConfigs()
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(config.Set, configCmd, proxy.Server)
		util.ExecCmds(commands, stdin)
	}

	closeFunc()
}
