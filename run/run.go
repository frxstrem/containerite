package run

import (
	"os/exec"
	"syscall"
)

func RunCommand(cmd *exec.Cmd) (int, error) {
	if err := cmd.Start(); err != nil {
		return -1, err
	}
	return WaitCommand(cmd)
}

func WaitCommand(cmd *exec.Cmd) (int, error) {
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			ws := exitErr.Sys().(syscall.WaitStatus)
			return ws.ExitStatus(), nil
		}
		return -1, err
	}
	return 0, nil
}
