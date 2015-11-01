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
		_, err := RemoteURL(tmpRoot)

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

		_, err := RemoteURL(tmpRoot)

		expected := "git remote is not defined"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestSourceUrl(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		sourceURL := url.SourceURL("master", "", 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestSourceUrlWithPath(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/master/" + path
		sourceURL := url.SourceURL("master", path, 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestSourceUrlWithPathAndRef(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/develop/" + path
		sourceURL := url.SourceURL("develop", path, 0, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestSourceUrlWithFrom(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30"
		sourceURL := url.SourceURL("develop", path, 30, 0)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestSourceUrlWithFromTo(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		path := "some/deep/dir.go"
		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/blob/develop/some/deep/dir.go#L30-L32"
		sourceURL := url.SourceURL("develop", path, 30, 32)
		if sourceURL != expected {
			t.Errorf(`got "%s" want "%s"`, sourceURL, expected)
		}
	})
}

func TestIssueUrl(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/issues"
		issueURL := url.IssueURL(0)
		if issueURL != expected {
			t.Errorf(`got "%s" want "%s"`, issueURL, expected)
		}
	})
}

func TestIssueUrlWithNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/issues/3"
		issueURL := url.IssueURL(3)
		if issueURL != expected {
			t.Errorf(`got "%s" want "%s"`, issueURL, expected)
		}
	})
}

func TestIssueUrlWithMinusNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		issueURL := url.IssueURL(-3)
		if issueURL != expected {
			t.Errorf(`got "%s" want "%s"`, issueURL, expected)
		}
	})
}

func TestPullrequestUrl(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/pulls"
		pullrequestURL := url.PullrequestURL(0)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}

func TestPullrequestUrlWithNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git/pulls/3"
		pullrequestURL := url.PullrequestURL(3)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}

func TestPullrequestUrlWithMinusNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeGirDir(tmpRoot)

		url, err := RemoteURL(tmpRoot)
		if err != nil {
			t.Errorf(`RemoteURL returned err "%s"`, err)
		}

		expected := "https://git.example.com/git/git"
		pullrequestURL := url.PullrequestURL(-3)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}
