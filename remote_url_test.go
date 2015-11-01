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
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot, "master")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		sourceURL := url.SourceURL("", 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestRemoteUrlWithPath(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "master")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/master/" + path
		sourceURL := url.SourceURL(path, 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestRemoteUrlWithPathAndRef(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/" + path
		sourceURL := url.SourceURL(path, 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestRemoteUrlWithFrom(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30"
		sourceURL := url.SourceURL(path, 30, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestRemoteUrlWithFromTo(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot, "develop")
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, path)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30-L32"
		sourceURL := url.SourceURL(path, 30, 32)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}
