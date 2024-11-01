//go:build windows

package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ExecCmds(commands []*Command) bool {
	var failed bool
	for _, command := range commands {
		var cmdStr string = command.Name

		// Combine command into single string if args are given
		if len(command.Arguments) > 0 {
			cmdArgs := strings.Join(command.Arguments, " ")
			cmdStr = fmt.Sprintf("%s %s", command.Name, cmdArgs)
		}

		// Try executing command in default shell
		cmd := exec.Command("powershell", "-NoLogo", "-NoProfile", "-Command", cmdStr)
		if err := cmd.Run(); err != nil {
			log.Printf("Command '%s' failed!\n", cmdStr)
			failed = true
		}
	}
	return failed
}
