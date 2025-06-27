package util

import "os/exec"

func execShell(command string) *exec.Cmd {
	return exec.Command("powershell", "-NoLogo", "-NoProfile", "-Command", command)
}
