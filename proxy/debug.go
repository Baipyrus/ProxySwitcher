package proxy

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func Debug(cfgPath string) {
	path, _ := filepath.Abs(cfgPath)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		path = "[N/A]"
	}

	fmt.Printf("\nConfig:\n")
	fmt.Printf("%s\n\n", path)

	proxy, err := readSystemProxy()
	proxyServer := proxy.Server
	if err != nil {
		log.Fatal(err)
	} else if proxyServer == "" {
		proxyServer = "[N/A]"
	}

	fmt.Println("System Proxy:")
	fmt.Printf("Enabled: %t\n", proxy.Enabled)
	fmt.Printf("Server: %s\n\n", proxyServer)

	configs, err := util.ReadConfigs(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, config := range configs {
		configCmd := config.Name
		// Use command instead of name, if given
		if config.Cmd != "" {
			configCmd = config.Cmd
		}
		fmt.Printf("Loading commands for '%s':\n", config.Name)

		// Debug Proxy Set Commands
		setCmds := generateCommands(configCmd, config.Set, "[PROXY PLACEHOLDER]")
		fmt.Println("Set Commands:")
		fmt.Printf("%s\n", mapCmdsToStr(setCmds))

		// Debug Proxy Unset Commands
		unsetCmds := generateCommands(configCmd, config.Unset, "")
		fmt.Println("Unset Commands:")
		fmt.Printf("%s\n\n", mapCmdsToStr(unsetCmds))
	}
}
