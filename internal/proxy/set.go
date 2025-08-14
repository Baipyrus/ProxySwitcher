package proxy

import (
	"log"

	"github.com/Baipyrus/ProxySwitcher/internal/util"
)

func Set(cfgPath string) {
	proxy, _ := ReadProxy(cfgPath)

	// Set system proxy and toggle
	setProxyState(true)
	proxy.Enabled = true

	var failed bool
	configs, _ := util.ReadConfigs(cfgPath, false)
	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}

		commands := generateCommands(configCmd, config.Set, proxy.Server)
		failed = util.ExecCmds(commands)
	}

	// Additional feedback on error
	if failed {
		log.Printf("One or more commands failed to execute. Run command 'debug' to see more.\n")
	}
}
