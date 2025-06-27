//go:build !windows

// Unimplemented "system proxy" functions for non-windows machines.
// TODO: Read GNOME Proxy or other Desktop Settings

package proxy

func setSystemProxy(_ *Proxy) error {
	return nil
}

func setProxyState(_ bool) error {
	return nil
}
