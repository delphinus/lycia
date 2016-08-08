package fssh

import (
	"bufio"
	"github.com/delphinus35/lycia/error"
	"os"
	"os/exec"
	"strings"
)

// Setting for FSSH
type Setting struct {
	port     string
	user     string
	copyArgs string
	path     string
}

func (setting Setting) sshArgs(cmd string) (result []string) {
	append(result,
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
	if fsshEnv, err := fetchFsshEnv(); err != nil {
		err = error.LyciaError(err)
		return
	} else if fsshEnv != empty {
		commandToExecute := strings.Join(append([]string{name}, args...), " ")
		name = "ssh"
		args = fsshEnv.sshArgs(commandToExecute)
	}
	cmd = exec.Command(name, args...)
}

func fetchFsshEnv() (result Setting, err error) {
	if tmuxEnv := os.Getenv("TMUX"); len(tmuxEnv) > 0 {
		result, err = parseTmuxEnv()
	} else {
		result = Setting{
			os.Getenv("LC_FSSH_PORT"),
			os.Getenv("LC_FSSH_USER"),
			os.Getenv("LC_FSSH_COPY_ARGS"),
			os.Getenv("LC_FSSH_PATH"),
		}
	}
	return
}

func parseTmuxEnv() (setting Setting, err error) {
	showenvCommand := exec.Command("tmux", "showenv")
	stdout, err = showenvCommand.StdoutPipe()
	if err != nil {
		err = error.LyciaError(err)
		return
	}

	err = showenvCommand.Start()
	if err != nil {
		err = error.LyciaError(err)
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
		keyNameIndex := strings.Index(keyAndValue[0], prefix)
		if keyNameIndex == -1 {
			return
		}
		switch keyAndValue[0][keyNameIndex:] {
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
		err = error.LyciaError(err)
		return
	}

	if len(port) > 0 && len(user) > 0 && len(copyArgs) > 0 && len(path) > 0 {
		setting = Setting{port, user, copyArgs, path}
	}
	return
}
