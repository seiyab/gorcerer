package process

import "os/exec"

func OverrideRunner(run func(name string, arg ...string) *exec.Cmd) {
	command = run
}

var command = exec.Command

func Command(name string, arg ...string) *exec.Cmd {
	return command(name, arg...)
}

func EchoRunner() func(name string, arg ...string) *exec.Cmd {
	return func(name string, arg ...string) *exec.Cmd {
		return exec.Command(
			"echo",
			append([]string{name}, arg...)...,
		)
	}
}
