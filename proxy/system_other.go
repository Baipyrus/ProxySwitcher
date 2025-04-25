//go:build !windows

// Unimplemented "system proxy" functions for non-windows machines.
// TODO: Read GNOME Proxy or other Desktop Settings

package proxy

func readSystemProxy() (*Proxy, error) {
	return &Proxy{}, nil
}

func setSystemProxy(_ bool) error {
	return nil
}
