package proxy

import (
	"errors"
	"strings"

	"golang.org/x/sys/windows/registry"
)

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
