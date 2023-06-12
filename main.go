package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"

	"github.com/pkg/errors"
)

func main() {
	// just create a process with a new set of namespaces
	var err error

	switch os.Args[1] {
	case "run":
		err = run(os.Args[2:]...)
	case "child":
		err = child(os.Args[2:]...)
	default:
		err = errors.New("invalid command")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func run(commands ...string) error {
	// Clone new namespaces for the child process
	// Execute the child process

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, commands...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET,
	}

	err := cmd.Run()

	if err != nil {
		return errors.Wrap(err, "couldn't run main process")
	}
	return nil
}

func child(commands ...string) error {
	var err error

	err = limit_memory(10)
	if err != nil {
		return errors.Wrap(err, "couldn't limit memory")
	}

	err = limit_pids(20)
	if err != nil {
		return errors.Wrap(err, "couldn't limit pids")
	}

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = syscall.Sethostname([]byte("gocker"))
	if err != nil {
		return errors.Wrap(err, "couldn't set hostname")
	}

	err = syscall.Chroot("/home/o/images/ubuntu")
	if err != nil {
		return errors.Wrap(err, "couldn't change root")
	}

	err = syscall.Chdir("/")
	if err != nil {
		return errors.Wrap(err, "couldn't change directory")
	}

	err = syscall.Mount("proc", "proc", "proc", 0, "")
	if err != nil {
		return errors.Wrap(err, "couldn't mount proc")
	}

	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "couldn't run child process")
	}

	err = syscall.Unmount("proc", 0)
	if err != nil {
		return errors.Wrap(err, "couldn't unmount proc")
	}

	return nil
}

func limit_memory(size int) error {
	var err error

	cgroups_path := "/sys/fs/cgroup"
	memory_path := path.Join(cgroups_path, "memory")
	new_group_path := path.Join(memory_path, "gocker")

	err = os.Mkdir(new_group_path, 0755)
	if err != nil && !os.IsExist(err) {
		return errors.Wrap(err, "couldn't create memory cgroup")
	}

	// write a value to the memory.limit_in_bytes file
	size_in_bytes := size * 1024 * 1024
	err = os.WriteFile(path.Join(new_group_path, "memory.limit_in_bytes"), []byte(fmt.Sprint(size_in_bytes)), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to memory.limit_in_bytes")
	}

	// clean up
	err = os.WriteFile(path.Join(new_group_path, "notify_on_release"), []byte("1"), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to notify_on_release")
	}

	// associate the process to the cgroup
	pid := strconv.Itoa(os.Getpid())
	err = os.WriteFile(path.Join(new_group_path, "cgroup.procs"), []byte(pid), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to cgroup.procs")
	}

	return nil
}

func limit_pids(num int) error {
	var err error

	// create new cgroup for pids
	cgroups_path := "/sys/fs/cgroup"
	pids_path := path.Join(cgroups_path, "pids")
	new_group_path := path.Join(pids_path, "gocker")

	err = os.Mkdir(new_group_path, 0755)
	if err != nil && !os.IsExist(err) {
		return errors.Wrap(err, "couldn't create pids cgroup")
	}

	// write a value to the pids.max file
	err = os.WriteFile(path.Join(new_group_path, "pids.max"), []byte(fmt.Sprint(num)), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to pids.max")
	}

	// clean up
	err = os.WriteFile(path.Join(new_group_path, "notify_on_release"), []byte("1"), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to notify_on_release")
	}

	// associate the process to the cgroup
	pid := strconv.Itoa(os.Getpid())
	err = os.WriteFile(path.Join(new_group_path, "cgroup.procs"), []byte(pid), 0700)
	if err != nil {
		return errors.Wrap(err, "couldn't write to cgroup.procs")
	}

	return nil
}
