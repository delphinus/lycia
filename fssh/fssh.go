package fssh

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/delphinus35/lycia/util"
)

var _os = struct {
	Getenv func(string) string
}{
	os.Getenv,
}

var _exec = struct {
	Command func(string, ...string) *exec.Cmd
}{
	exec.Command,
}

// Setting for FSSH
type Setting struct {
	port     string
	user     string
	copyArgs string
	path     string
}

func (setting Setting) sshArgs(cmd string) (result []string) {
	result = append(result,
		"-p", setting.port,
		"-l", setting.user,
		setting.copyArgs,
		"localhost",
		"PATH="+setting.path,
		cmd,
	)
	return
}

// Command is to execute any commands with FSSH
func Command(name string, args ...string) (cmd *exec.Cmd, err error) {
	empty := Setting{}
	var fsshEnv Setting
	if fsshEnv, err = fetchFsshEnv(); err != nil {
		err = util.LyciaError(err.Error())
		return
	} else if fsshEnv != empty {
		commandToExecute := strings.Join(append([]string{name}, args...), " ")
		name = "ssh"
		args = fsshEnv.sshArgs(commandToExecute)
	}
	cmd = _exec.Command(name, args...)
	return
}

func fetchFsshEnv() (result Setting, err error) {
	if tmuxEnv := _os.Getenv("TMUX"); len(tmuxEnv) > 0 {
		result, err = parseTmuxEnv()
	} else {
		result = Setting{
			_os.Getenv("LC_FSSH_PORT"),
			_os.Getenv("LC_FSSH_USER"),
			_os.Getenv("LC_FSSH_COPY_ARGS"),
			_os.Getenv("LC_FSSH_PATH"),
		}
	}
	return
}

func parseTmuxEnv() (setting Setting, err error) {
	showenvCommand := _exec.Command("tmux", "showenv")
	stdout, err := showenvCommand.StdoutPipe()
	if err != nil {
		err = util.LyciaError(err.Error())
		return
	}

	err = showenvCommand.Start()
	if err != nil {
		err = util.LyciaError(err.Error())
		return
	}

	scanner := bufio.NewScanner(stdout)
	const sep = "="
	const prefix = "LC_FSSH_"
	var port, user, copyArgs, path string
	for scanner.Scan() {
		keyAndValue := strings.SplitN(scanner.Text(), sep, 2)
		if len(keyAndValue) != 2 {
			return
		}
		switch keyAndValue[0][len(prefix):] {
		case "PORT":
			port = keyAndValue[1]
		case "USER":
			user = keyAndValue[1]
		case "COPY_ARGS":
			copyArgs = keyAndValue[1]
		case "PATH":
			path = keyAndValue[1]
		}
	}

	err = showenvCommand.Wait()
	if err != nil {
		err = util.LyciaError(err.Error())
		return
	}

	if len(port) > 0 && len(user) > 0 && len(copyArgs) > 0 && len(path) > 0 {
		setting = Setting{port, user, copyArgs, path}
	}
	return
}
