//go:build windows

package util

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Create a single powershell process and leave closing, input and output open
func ReadyCmd() (*io.WriteCloser, func() error, error) {
	cmd := exec.Command("powershell", "-NoLogo", "-NoProfile", "-Command", "-")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	return &stdin, func() error {
		// Close stdin pipe
		stdin.Close()

		// Wait for command to flush
		err := cmd.Wait()
		return err
	}, nil
}

func ExecCmds(commands []*Command, stdin *io.WriteCloser) {
	for _, command := range commands {
		cmdArgs := strings.Join(command.Arguments, " ")
		cmdStr := fmt.Sprintf("%s %s", command.Name, cmdArgs)
		fmt.Fprintln(*stdin, cmdStr)
	}
}
