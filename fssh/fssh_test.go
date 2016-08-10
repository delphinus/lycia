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
		SOME_INVALID_ENV=hoge
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
		LC_FSSH_PORT=3333
		LC_FSSH_USER=delphinus
		LC_FSSH_COPY_ARGS=hoge fuga=hogefuga
		LC_FSSH_PATH=/tmp/test/test
	`, func() {
		setting, _ := parseTmuxEnv()
		expected := Setting{
			"3333",
			"delphinus",
			"hoge fuga=hogefuga",
			"/tmp/test/test",
		}
		if setting != expected {
			t.Errorf(`get "%v" want "%v"`, setting, expected)
		}
	})
}
