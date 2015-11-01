package fssh

import (
	"os"
	"os/exec"
	"strings"
)

var fsshENV = map[string]string{
	"port":      os.Getenv("LC_FSSH_PORT"),
	"user":      os.Getenv("LC_FSSH_USER"),
	"copy_args": os.Getenv("LC_FSSH_COPY_ARGS"),
	"path":      os.Getenv("LC_FSSH_PATH"),
}
var IsFssh = fsshENV["port"] != ""

func Command(name string, args ...string) *exec.Cmd {
	if IsFssh {
		commandToExecute := strings.Join(append([]string{name}, args...), " ")
		name = "ssh"
		args = []string{
			"-p", fsshENV["port"],
			"-l", fsshENV["user"],
			fsshENV["copy_args"],
			"localhost",
			"PATH=" + fsshENV["path"],
			commandToExecute,
		}
	}
	return exec.Command(name, args...)
}
