package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/frxstrem/containerite/options"
	"github.com/frxstrem/containerite/run"
	"github.com/frxstrem/containerite/slave"
)

func main() {
	var err error

	opts := options.Options{}
	err = opts.Parse(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	var exitCode int
	if opts.IsSlave {
		exitCode, err = slave.Run(opts)
	} else {
		exitCode, err = spawnSlave(os.Args[1:])
	}

	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(exitCode)
	return
}

func spawnSlave(args []string) (int, error) {
	args = append([]string{"--x-slave"}, args...)

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS,
	}

	return run.RunCommand(cmd)
}
