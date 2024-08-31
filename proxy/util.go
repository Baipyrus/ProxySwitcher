package proxy

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

func ReadSystemProxy() (*Proxy, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	enabled, _, err := key.GetIntegerValue("proxyEnable")
	if err != nil {
		return nil, err
	}

	server, _, err := key.GetStringValue("proxyServer")
	if err != nil && !errors.Is(err, registry.ErrNotExist) {
		return nil, err
	}

	return &Proxy{Enabled: enabled != 0, Server: server}, nil
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
	if err != nil {
		return err
	}

	return nil
}
