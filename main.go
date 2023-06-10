package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

/*
	it will create a /bin/sh process in a new set of namespaces
*/

func main() {
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

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, commands[0:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWNS |
	// 		syscall.CLONE_NEWUTS |
	// 		syscall.CLONE_NEWIPC |
	// 		syscall.CLONE_NEWPID |
	// 		syscall.CLONE_NEWNET |
	// 		syscall.CLONE_NEWUSER,
	// }

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER,
	}

	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func child(commands ...string) error {

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PS1=-[gocker]- # "}

	err := cmd.Run()

	if err != nil {
		return errors.New("couldn't child")
	}
	return nil
}
