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

func applyProxy(configArgs []string, configCmd string, proxyServer string, variant *util.Variant) ([]string, string) {
	// Skip, no proxy provided
	if proxyServer == "" {
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
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	enableVal, _, err := key.GetIntegerValue("proxyEnable")
	if err != nil {
		return nil, err
	}
	enabled := enableVal != 0

	servers, _, err := key.GetStringValue("proxyServer")
	if err != nil && !errors.Is(err, registry.ErrNotExist) {
		return nil, err
	}

	if !strings.ContainsAny(servers, ";=") {
		return &Proxy{Enabled: enabled, Server: servers}, nil
	}

	serverSplit := strings.Split(servers, ";")
	serverDict := make(map[string]string)
	for _, substr := range serverSplit {
		subSplit := strings.Split(substr, "=")
		key, value := subSplit[0], subSplit[1]
		serverDict[key] = value
	}

	if serverDict["http"] != "" {
		return &Proxy{Enabled: enabled, Server: serverDict["http"]}, nil
	}

	if serverDict["https"] != "" {
		return &Proxy{Enabled: enabled, Server: serverDict["https"]}, nil
	}

	return nil, errors.New("You need configure either HTTP or HTTPS proxy to proceed.")
}

func SetSystemProxy(state bool) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	var value uint32
	if state {
		value = 1
	} else {
		value = 0
	}

	err = key.SetDWordValue("proxyEnable", value)
	return err
}
