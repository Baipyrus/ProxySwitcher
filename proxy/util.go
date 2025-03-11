//go:build windows

package proxy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
	"golang.org/x/sys/windows/registry"
)

func processVars(cmd *util.Command, isVariableType bool) {
	// If not a variable type, return early as there's nothing to replace
	if !isVariableType {
		return
	}

	// Replace specific $PRSW_ARG in the command's Name with each argument
	for _, arg := range cmd.Arguments {
		cmd.Name = strings.Replace(cmd.Name, "$PRSW_ARG", arg, 1)
	}

	cmd.Arguments = nil
}

func injectProxy(cmd *util.Command, variant *util.Variant, proxyServer string) {
	// Skip if no proxy is provided or proxy option is discarded
	if proxyServer == "" || variant.DiscardProxy {
		return
	}

	// Surround 'proxyServer' with the provided surround string, if any
	if variant.Surround != "" {
		proxyServer = fmt.Sprintf("%[1]s%[2]s%[1]s", variant.Surround, proxyServer)
	}

	// Insert proxy into last place of command if the variant is VARIABLE
	if variant.Type == util.VARIABLE && strings.Count(cmd.Name, "$PRSW_ARG") == 1 {
		cmd.Name = strings.Replace(cmd.Name, "$PRSW_ARG", proxyServer, 1)
		return
	}

	// Insert proxy after the equator if specified
	if variant.Equator != "" {
		lastArgIdx := len(cmd.Arguments) - 1
		cmd.Arguments[lastArgIdx] += variant.Equator + proxyServer
		return
	}

	// Otherwise, append the proxy as a new argument
	cmd.Arguments = append(cmd.Arguments, proxyServer)
}

func generateCommands(base string, variants []*util.Variant, proxyServer string) []*util.Command {
	var commands []*util.Command

	// Iterate through all variants and generate a command for each
	for _, variant := range variants {
		isVariableType := variant.Type == util.VARIABLE

		// Create command from default parameters
		cmd := &util.Command{
			Name:      base,
			Arguments: append([]string{}, variant.Arguments...),
		}

		processVars(cmd, isVariableType)
		injectProxy(cmd, variant, proxyServer)

		commands = append(commands, cmd)
	}

	return commands
}

func readSystemProxy() (*Proxy, error) {
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

func setSystemProxy(state bool) error {
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
