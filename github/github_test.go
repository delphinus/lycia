package github

import (
	"net/url"
	"regexp"
	"testing"
)

var repoURL, _ = url.Parse("https://github.com/delphinus35/lycia")
var repo = Repository(repoURL)
var hogeURL, _ = url.Parse("https://example.com/hogehoge")
var hogeRepo = Repository(hogeURL)

func TestPullrequestUrlWithNumber(t *testing.T) {
	expected := repoURL.String() + "/pulls"
	pullrequestURL := repo.PullrequestUrlWithNumber(0)
	if pullrequestURL != expected {
		t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
	}
}

func TestPullrequestUrlWithNumberWithNum(t *testing.T) {
	expected := repoURL.String() + "/pull/3"
	pullrequestURL := repo.PullrequestUrlWithNumber(3)
	if pullrequestURL != expected {
		t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
	}
}

func TestPullrequestUrlWithNumberWithMinusNum(t *testing.T) {
	expected := repoURL.String()
	pullrequestURL := repo.PullrequestUrlWithNumber(-3)
	if pullrequestURL != expected {
		t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
	}
}

func TestPullrequestUrlWithBranch(t *testing.T) {
	_, err := repo.PullrequestUrlWithBranch("hoge")
	expected := "pullrequest not found for the branch: hoge"
	if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
		t.Errorf(`got "%s" want "%s"`, err, expected)
	}

	_, err = repo.PullrequestUrlWithBranch("feature/detect-pr-for-branch-at-present")
	if err != nil {
		t.Errorf(`err was found: "%s"`, err.Error())
	}

	/*
		_, err = hogeRepo.PullrequestUrlWithBranch("hoge")
		expected = "failed to fetch:"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	*/
}
