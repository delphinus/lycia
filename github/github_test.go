package github

import (
	"net/url"
	"regexp"
	"testing"
)

var repoURL, _ = url.Parse("https://github.com/delphinus35/lycia")
var repo, _ = Repository(repoURL)
var hogeURL, _ = url.Parse("https://example.com/hogehoge")
var hogeRepo, _ = Repository(hogeURL)

func TestPullrequestUrlWithNumber(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		expected := repoURL.String() + "/pulls"
		pullrequestURL := repo.PullrequestUrlWithNumber(0)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}

func TestPullrequestUrlWithNumberWithNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		expected := repoURL.String() + "/pull/3"
		pullrequestURL := repo.PullrequestUrlWithNumber(3)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}

func TestPullrequestUrlWithNumberWithMinusNum(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		expected := repoURL.String()
		pullrequestURL := repo.PullrequestUrlWithNumber(-3)
		if pullrequestURL != expected {
			t.Errorf(`got "%s" want "%s"`, pullrequestURL, expected)
		}
	})
}

func TestPullrequestUrlWithBranch(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		_, err := repo.PullrequestUrlWithBranch("hoge", false)
		expected := "pullrequest not found for the branch: hoge"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}

		_, err = repo.PullrequestUrlWithBranch("feature/detect-pr-for-branch-at-present", false)
		if err != nil {
			t.Errorf(`err was found: "%s"`, err.Error())
		}

		err = repo.Cache.LoadCache()
		if err != nil {
			t.Errorf(`err was found: "%s"`, err.Error())
		}

		if _, ok := repo.Cache["https://github.com/delphinus35/lycia"]; !ok {
			t.Errorf(`branchToUrl map is not found in cache`)
		}

		if _, ok := repo.Cache["https://github.com/delphinus35/lycia"]["feature/detect-pr-for-branch-at-present"]; !ok {
			t.Errorf(`prUrl is not found in branchToUrl map`)
		}
	})
}
