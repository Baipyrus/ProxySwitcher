package proxy

import (
	"golang.org/x/sys/windows/registry"
)

func setSystemProxy(p *Proxy) error {
	// Open registry key for internet settings
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	// Get state as int instead of bool
	var state uint32
	if p.Enabled {
		state = 1
	} else {
		state = 0
	}

	// Set enable proxy according to boolean
	err = key.SetDWordValue("proxyEnable", state)
	if err != nil {
		return err
	}

	// Set proxy server URL
	return key.SetStringValue("proxyServer", p.Server)
}

func setProxyState(state bool) error {
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
