package proxy

import (
	"log"

	"github.com/Baipyrus/ProxySwitcher/util"
)

func Unset(cfgPath string) {
	proxy, _ := readSystemProxy()
	// Unset system proxy, if not already
	if proxy.Enabled {
		setSystemProxy(false)
	}

	var failed bool
	configs, _ := util.ReadConfigs(cfgPath)
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(configCmd, config.Unset, "")
		failed = util.ExecCmds(commands)
	}

	// Additional feedback on error
	if failed {
		log.Printf("One or more commands failed to execute. Run command 'debug' to see more.\n")
	}
}
