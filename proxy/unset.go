//go:build windows

package proxy

import (
	"github.com/Baipyrus/ProxySwitcher/util"
)

func Unset(cfgPath string) {
	stdin, closeFunc, _ := util.ReadyCmd()

	proxy, _ := ReadSystemProxy()
	// Unset system proxy, if not already
	if proxy.Enabled {
		SetSystemProxy(false)
	}

	configs, _ := util.ReadConfigs(cfgPath)
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(configCmd, config.Unset, "")
		util.ExecCmds(commands, stdin)
	}

	closeFunc()
}
