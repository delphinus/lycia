package github

import (
	"encoding/json"
	"fmt"
	. "github.com/delphinus35/lycia/error"
	"net/http"
	"net/url"
	"strings"
)

type repository struct {
	URL *url.URL
}

func Repository(repoURL *url.URL) *repository {
	return &repository{repoURL}
}

func (repo *repository) PullrequestUrlWithNumber(num int) (pullrequestURL string) {
	if num == 0 {
		pullrequestURL = repo.URL.String() + "/pulls"
	} else if num > 0 {
		pullrequestURL = fmt.Sprintf("%s/pull/%d", repo.URL.String(), num)
	} else {
		pullrequestURL = repo.URL.String()
	}
	return
}

func (repo *repository) PullrequestUrlWithBranch(branch string) (prURL *url.URL, err error) {
	values := url.Values{}
	repoPath := strings.TrimLeft(repo.URL.Path, "/")
	queryString := fmt.Sprintf("repo:%s type:pr head:%s", repoPath, branch)
	values.Add("q", queryString)

	searchURL := fmt.Sprintf("https://api.github.com/search/issues?%s", values.Encode())
	res, err := http.Get(searchURL)
	if err != nil {
		err = LyciaError("failed to fetch: " + searchURL)
		return
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var searchIssues SearchIssues
	if err = decoder.Decode(&searchIssues); err != nil {
		return
	}

	items := searchIssues.Items
	if len(items) == 0 {
		err = LyciaError(fmt.Sprintf("pullrequest not found for the branch: %s", branch))
		return
	}

	prURL, err = url.Parse(items[0].HtmlUrl)
	if err != nil {
		err = LyciaError(fmt.Sprintf("html_url is invalid: %s", items[0].HtmlUrl))
	}
	return
}
