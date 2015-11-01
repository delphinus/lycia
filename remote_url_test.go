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

func withFakeGirDir(tmpRoot string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpRoot
	_ = cmd.Run()
	cmd = exec.Command("git", "remote", "add", "origin", "git://git.example.com/git/git")
	cmd.Dir = tmpRoot
	_ = cmd.Run()
}

func TestRemoteUrlError(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		_, err := RemoteURL(tmpRoot, "some_ref", "", 0, 0)

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

		_, err := RemoteURL(tmpRoot, "master", "", 0, 0)

		expected := "git remote is not defined"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestRemoteUrlWithRemote(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot, "master", "", 0, 0)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		if url != nil && url.String() != expected {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}

func TestRemoteUrlWithPath(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "master", path, 0, 0)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/master/" + path
		if url != nil && url.String() != expected {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}

func TestRemoteUrlWithPathAndRef(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop", path, 0, 0)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/" + path
		if url != nil && url.String() != expected {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}

func TestRemoteUrlWithFrom(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop", path, 30, 0)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30"
		if url != nil && url.String() != expected {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}

func TestRemoteUrlWithFromTo(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop", path, 30, 32)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30-32"
		if url != nil && url.String() != expected {
			t.Errorf(`got "%s" want "%s"`, url, expected)
		}
	})
}
