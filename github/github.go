package github

import (
	"fmt"
	"net/http"
	"net/url"
)

type repository struct {
	URL url.URL
}

func Repository(repoURL url.URL) *repository {
	return &repository{repoURL}
}

func (repo *repository) PullrequestURL(branch string) (prURL *url.URL, err error) {
	values := url.Values{}
	queryString := fmt.Sprintf("repo:%s type:pr is:open head:%s", path, branch)
	values.Add("q", queryString)

	searchURL := fmt.Sprintf("%s/search/issues?%s", repo.URL, values.Encode())
	res, err := http.Get(searchURL)
	if err != nil {
		err = GitUrlError("failed to fetch: " + searchURL)
		return
	}

	defer res.Body.Close()
	return
}
