package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

func withFakeEnv(t *testing.T, block func(string)) {
	tmpRoot, err := ioutil.TempDir("", "lycia_")
	if err != nil {
		t.Fatalf("cloud not create tempdir: %s", err)
	}
	defer func() { os.RemoveAll(tmpRoot) }()

	block(tmpRoot)
}

func TestRemoteUrlError(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		_, err := RemoteURL(tmpRoot, "some_ref")

		expected := "can not exec 'git remove -v'"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestRemoteUrlNoRemote(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		cmd := exec.Command("git", "init")
		cmd.Dir = tmpRoot
		_ = cmd.Run()

		_, err := RemoteURL(tmpRoot, "master")

		expected := "git remote is not defined"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestRemoteUrlWithRemote(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		cmd := exec.Command("git", "init")
		cmd.Dir = tmpRoot
		_ = cmd.Run()
		cmd = exec.Command("git", "remote", "add", "origin", "git://git.example.com/git/git")
		cmd.Dir = tmpRoot
		_ = cmd.Run()

		url, err := RemoteURL(tmpRoot, "master")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		if url != nil && url.String() != "https://git.example.com/git/git" {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}
