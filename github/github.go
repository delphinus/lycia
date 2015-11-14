package github

import (
	"encoding/json"
	"fmt"
	. "github.com/delphinus35/lycia/error"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type repository struct {
	URL url.URL
}

func Repository(repoURL url.URL) *repository {
	return &repository{repoURL}
}

func (repo *repository) PullrequestURL(branch string) (prURL *url.URL, err error) {
	values := url.Values{}
	repoPath := strings.TrimLeft(repo.URL.Path, "/")
	queryString := fmt.Sprintf("repo:%s type:pr is:open head:%s", repoPath, branch)
	values.Add("q", queryString)

	searchURL := fmt.Sprintf("https://api.github.com/search/issues?%s", values.Encode())
	res, err := http.Get(searchURL)
	if err != nil {
		err = LyciaError("failed to fetch: " + searchURL)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = LyciaError("failed to read body: " + searchURL)
		return
	}

	var searchIssue SearchIssue
	if err = json.Unmarshal(body, &searchIssue); err != nil {
		return
	}

	items := searchIssue.Items
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
