package proxy

import (
	"github.com/Baipyrus/ProxySwitcher/util"
)

func Unset(cfgFile string) {
	stdin, closeFunc, _ := util.ReadyCmd()

	proxy, _ := ReadSystemProxy()
	// Unset system proxy, if not already
	if proxy.Enabled {
		SetSystemProxy(false)
	}

	configs, _ := util.ReadConfigs(cfgFile)
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
