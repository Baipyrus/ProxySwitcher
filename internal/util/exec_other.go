//go:build !windows

package util

import "os/exec"

func execShell(command string) *exec.Cmd {
	return exec.Command("bash", "--noprofile", "--norc", "-c", command)
}
