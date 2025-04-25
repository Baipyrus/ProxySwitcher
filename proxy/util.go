package proxy

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
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
