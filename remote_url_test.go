package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// testing exec.Command
// referred from http://npf.io/2015/06/testing-exec-command/
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestError(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	_, err := RemoteURL(".", "some_ref")
	expected := "can not exec 'git remove -v' : exit status 2"
	if err.Error() != expected {
		t.Errorf(errFormat, err, expected)
	}
}

// TestHelperProcess isn't a real test. It's used as a helper process
// for TestParameterRun.
func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	switch cmd {
	case "git":
		iargs := []interface{}{}
		for _, s := range args {
			iargs = append(iargs, s)
		}
		//fmt.Println(iargs...)
		fmt.Fprintf(os.Stderr, "HOGE error\n")
		os.Exit(2)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}
