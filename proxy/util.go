package proxy

import (
	"errors"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
	"golang.org/x/sys/windows/registry"
)

func readArgs(replaceVariable bool, args []string, configCmd string) ([]string, string) {
	var configArgs []string

	for _, arg := range args {
		// If not replacable, append
		if !replaceVariable {
			configArgs = append(configArgs, arg)
			continue
		}

		// Replace specific "ProxySwitcher Argument" in command
		configCmd = strings.Replace(configCmd, "$PRSW_ARG", arg, 1)
	}

	return configArgs, configCmd
}

func applyProxy(configArgs []string, configCmd, proxyServer string, variant *util.Variant) ([]string, string) {
	// Skip, no proxy provided or proxy option discarded
	if proxyServer == "" || variant.DiscardProxy {
		return configArgs, configCmd
	}

	// Insert proxy only on last VARIABLE type
	if variant.Type == util.VARIABLE && strings.Count(configCmd, "$PRSW_ARG") == 1 {
		configCmd = strings.Replace(configCmd, "$PRSW_ARG", proxyServer, 1)
		return configArgs, configCmd
	}

	// Insert proxy right after equator
	if variant.Equator != "" {
		configArgs[len(configArgs)-1] += variant.Equator + proxyServer
		return configArgs, configCmd
	}

	// Or otherwise just append it as an argument
	configArgs = append(configArgs, proxyServer)

	return configArgs, configCmd
}

func generateCommands(variants []*util.Variant, configCmd, proxyServer string) []*util.Command {
	var commands []*util.Command

	// Generate one command per variant
	for _, variant := range variants {
		replaceVariable := variant.Type == util.VARIABLE

		configArgs, configCmd := readArgs(replaceVariable, variant.Arguments, configCmd)
		configArgs, configCmd = applyProxy(configArgs, configCmd, proxyServer, variant)

		commands = append(commands, &util.Command{Name: configCmd, Arguments: configArgs})
	}

	return commands
}

func ReadSystemProxy() (*Proxy, error) {
	// Open registry key for internet settings
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	// Read registry value for proxy enabled
	enableVal, _, err := key.GetIntegerValue("proxyEnable")
	if err != nil {
		return nil, err
	}
	// Convert int value to bool
	enabled := enableVal != 0

	// Read registry value for proxy servers
	servers, _, err := key.GetStringValue("proxyServer")
	if err != nil && !errors.Is(err, registry.ErrNotExist) {
		return nil, err
	}

	// Use entire value if singular server
	proxy := &Proxy{Enabled: enabled, Server: ""}
	if !strings.ContainsAny(servers, ";=") {
		proxy.Server = servers
		return proxy, nil
	}

	// Map proxy servers into dictionary
	serverSplit := strings.Split(servers, ";")
	serverDict := make(map[string]string)
	for _, substr := range serverSplit {
		subSplit := strings.Split(substr, "=")
		key, value := subSplit[0], subSplit[1]
		serverDict[key] = value
	}

	// Grab HTTP proxy server first
	if serverDict["http"] != "" {
		proxy.Server = serverDict["http"]
		return proxy, nil
	}

	// Grab HTTP proxy server second
	if serverDict["https"] != "" {
		proxy.Server = serverDict["https"]
		return proxy, nil
	}

	// Return with empty proxy server ("discarded"; not detected)
	return proxy, nil
}

func SetSystemProxy(state bool) error {
	// Open registry key for internet settings
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	// Get state as int instead of bool
	var value uint32
	if state {
		value = 1
	} else {
		value = 0
	}

	// Write registry value to enable/disable proxy
	err = key.SetDWordValue("proxyEnable", value)
	return err
}
