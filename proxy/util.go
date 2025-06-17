package proxy

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
	"gopkg.in/ini.v1"
)

func SplitProxyUrl(url string) (string, string, string) {
	re := regexp.MustCompile(`^(\w+)://([^:/]+)(?::(\d+))?`)
	matches := re.FindStringSubmatch(url)

	protocol := matches[1]
	host := matches[2]

	var port string
	if len(matches) == 4 {
		port = matches[3]
	} else {
		switch protocol {
		case "http":
			port = "80"
		case "https":
			port = "443"
		}
	}

	return protocol, host, port
}

func SaveProxy(cfgPath string, p *Proxy) {
	cfg := ini.Empty()
	section, _ := cfg.NewSection("")

	protocol, host, port := SplitProxyUrl(p.Server)

	section.NewKey("protocol", protocol)
	section.NewKey("host", host)
	section.NewKey("port", port)

	proxyConf := path.Join(cfgPath, "proxy.conf")
	cfg.SaveTo(proxyConf)
}

func allKeysExist[T comparable](have []T, want []T) bool {
	for _, k := range want {
		if !slices.Contains(have, k) {
			return false
		}
	}
	return true
}

func ReadProxy(cfgPath string) (*Proxy, error) {
	proxy := &Proxy{}

	proxyConf := path.Join(cfgPath, "proxy.conf")
	cfg, err := ini.Load(proxyConf)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("")
	keys := section.KeyStrings()
	if !allKeysExist(keys, []string{"protocol", "host", "port"}) {
		return nil, errors.New("Proxy configuration is missing required entries!")
	}

	proxy.Server = fmt.Sprintf(
		"%s://%s:%s",
		section.Key("protocol").String(),
		section.Key("host").String(),
		section.Key("port").String(),
	)

	return proxy, nil
}

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
