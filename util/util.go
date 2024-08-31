package util

import (
	"bytes"
	"io"
	"os/exec"
)

func ReadyCmd() (*io.WriteCloser, func() error, error) {
	cmd := exec.Command("powershell", "-NoLogo", "-NoProfile", "-Command", "-")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	return &stdin, func() error {
		stdin.Close()

		if err := cmd.Wait(); err != nil {
			return err
		}

		return nil
	}, nil
}
