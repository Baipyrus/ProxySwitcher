package proxy

import (
	"github.com/Baipyrus/ProxySwitcher/util"
)

func Set(cfgPath string) {
	stdin, closeFunc, _ := util.ReadyCmd()

	proxy, _ := ReadSystemProxy()
	// Set system proxy, if not already
	if !proxy.Enabled {
		SetSystemProxy(true)
	}

	configs, _ := util.ReadConfigs(cfgPath)
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(configCmd, config.Set, proxy.Server)
		util.ExecCmds(commands, stdin)
	}

	closeFunc()
}
