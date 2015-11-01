package fssh

import (
	"os"
	"os/exec"
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
		args = append([]string{
			"-p", fsshENV["LC_FSSH_PORT"],
			"-l", fsshENV["LC_FSSH_USER"],
			fsshENV["LC_FSSH_COPY_ARGS"],
			"localhost",
			"PATH=" + fsshENV["LC_FSSH_PATH"],
			name,
		}, args...)
		name = "ssh"
	}
	return exec.Command(name, args...)
}
