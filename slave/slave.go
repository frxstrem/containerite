package slave

import (
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/frxstrem/containerite/run"

	"github.com/frxstrem/containerite/ex"
	"github.com/frxstrem/containerite/options"
)

func Run(opts options.Options) (int, error) {
	//defer ex.Autorecover(&err)

	initMountNamespace(opts)

	// run command
	cmd := exec.Command(opts.Command, opts.Arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return run.RunCommand(cmd)
}

func initMountNamespace(opts options.Options) {
	oldroot := ".oldroot"
	putold := path.Join(opts.Root, oldroot)

	// make old root private
	ex.Must(syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""))

	// create self-bind mount for target
	ex.Must(syscall.Mount(opts.Root, opts.Root, "", syscall.MS_BIND, ""))

	// create directory for putold
	ex.Must(os.MkdirAll(putold, 0700))

	// do root pivot
	ex.Must(syscall.PivotRoot(opts.Root, putold))
	ex.Must(os.Chdir("/"))

	// remove old root
	ex.Must(syscall.Unmount(oldroot, syscall.MNT_DETACH))
  ex.Must(os.Remove(oldroot))

  // mount kernel mountpoints
  ex.Must(syscall.Mount("proc", "/proc", "proc", 0, ""))
  ex.Must(syscall.Mount("sys", "/sys", "sysfs", 0, ""))
  ex.Must(syscall.Mount("devpts", "/dev/pts", "devpts", 0, ""))
}
