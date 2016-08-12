package fssh

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/MakeNowJust/heredoc"
)

func withFakeCommand(t *testing.T, envValues string, block func()) {
	_exec.Command = func(name string, arg ...string) *exec.Cmd {
		valueFile, _ := ioutil.TempFile("", "valueFile")
		valueFile.WriteString(heredoc.Doc(envValues))
		valueFile.Close()
		dummy, _ := ioutil.TempFile("", "dummy")
		dummy.WriteString(heredoc.Docf(`
			#!/bin/sh
			cat %s
		`, valueFile.Name()))
		dummy.Close()
		t.Logf("dummy: %s", dummy.Name())
		os.Chmod(dummy.Name(), 0755)
		return exec.Command(dummy.Name())
	}
	block()
}

func TestInvalidEnv(t *testing.T) {
	withFakeCommand(t, `
		HOGE=hoge
		FUGA
		LC_FSSH_HOGE=hogehoge
	`, func() {
		setting, _ := parseTmuxEnv()
		expected := Setting{}
		if setting != expected {
			t.Errorf(`get "%v" want "%v"`, setting, expected)
		}
	})
}

func TestValidEnv(t *testing.T) {
	withFakeCommand(t, `
		-DISPLAY
		LC_FSSH_COPY=reattach-to-user-namespace pbcopy
		LC_FSSH_COPY_ARGS=-o HostKeyAlias=hogehoge.local
		LC_FSSH_PASTE=reattach-to-user-namespace pbpaste
		LC_FSSH_PASTE_ARGS=-o HostKeyAlias=hogehoge.locall
		LC_FSSH_PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
		LC_FSSH_PORT=3333
		LC_FSSH_USER=delphinus
		-SSH_AGENT_PID
		SSH_ASKPASS=/usr/libexec/openssh/gnome-ssh-askpass
		SSH_AUTH_SOCK=/home/game/.ssh/auth_sock
		SSH_CONNECTION=192.168.11.11 12345 192.168.11.22 22
		-WINDOWID
		-XAUTHORITY
	`, func() {
		setting, _ := parseTmuxEnv()
		expected := Setting{
			"3333",
			"delphinus",
			"-o HostKeyAlias=hogehoge.local",
			"/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
		}
		if setting != expected {
			t.Errorf(`get "%v" want "%v"`, setting, expected)
		}
	})
}
