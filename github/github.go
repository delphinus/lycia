package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/delphinus35/lycia/ask"
	. "github.com/delphinus35/lycia/error"
	"net/http"
	"net/url"
	"strings"
)

type repository struct {
	URL    *url.URL
	Config Config
}

func Repository(repoURL *url.URL) *repository {
	config := make(Config)
	config.LoadConfig()
	return &repository{repoURL, config}
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

	apiRoot, err := repo.DetectApiRootAndSetAccessToken(values)
	if err != nil {
		err = LyciaError("cannot detect ApiRoot: " + err.Error())
		return
	}

	searchURL := fmt.Sprintf("%s/search/issues?%s", apiRoot, values.Encode())
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

func (repo *repository) NewAccessToken(host string, apiRoot string) (accessToken string, err error) {
	username, err := Ask(fmt.Sprintf("Please input username for '%s':", host), false)
	if err != nil {
		return
	}
	password, err := Ask(fmt.Sprintf("Please input password for '%s':", host), true)
	if err != nil {
		return
	}

	reqBody := `{"scopes":["repo"],"note":"lycia"}`
	req, err := repo.NewPostRequest(username, password, apiRoot+"/authorizations", reqBody)
	if err != nil {
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode == 401 {
		if v := res.Header.Get("X-Github-Otp"); strings.HasPrefix(v, "required;") {
			var otp string
			otp, err = Ask("Please input two-factor authentication code:", false)
			if err != nil {
				return
			}

			req, err = repo.NewPostRequest(username, password, apiRoot+"/authorizations", reqBody)
			if err != nil {
				return
			}
			req.Header.Add("X-Github-Otp", otp)
			res, err = client.Do(req)
			if err != nil {
				return
			}

		} else {
			err = LyciaError("Bad credentials")
			return
		}

	} else if res.StatusCode != 200 {
		err = LyciaError(fmt.Sprintf("Unknown status: %s", res.Status))
		return
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var authorizations Authorizations
	if err = decoder.Decode(&authorizations); err != nil {
		return
	}

	accessToken = authorizations.HashedToken
	if accessToken == "" {
		err = LyciaError("cannot detect HashedToken")
	}
	return
}

func (repo *repository) DetectApiRootAndSetAccessToken(values url.Values) (apiRoot string, err error) {
	if repo.URL.Host == "github.com" {
		apiRoot = "https://api.github.com"
		return
	}

	if v, ok := repo.Config[repo.URL.Host]; ok {
		apiRoot = v.ApiRoot
		values.Add("access_token", v.AccessToken)
		return
	}

	sc := SiteConfig{repo.URL.Host, "", ""}

	msg := fmt.Sprintf("Please input github API root path for '%s' (such as 'https://api.github.com') :", repo.URL.Host)
	apiRoot, err = Ask(msg, false)
	if err != nil {
		return
	}
	sc.ApiRoot = strings.TrimLeft(apiRoot, "/")

	sc.AccessToken, err = repo.NewAccessToken(sc.Host, sc.ApiRoot)
	if err != nil {
		err = LyciaError("failed to generate new access token: " + err.Error())
		return
	}

	repo.Config[repo.URL.Host] = sc
	err = repo.Config.SaveConfig()
	if err != nil {
		err = LyciaError("failed to save config: " + err.Error())
		return
	}

	values.Add("access_token", sc.AccessToken)
	return
}

func (repo *repository) NewPostRequest(username string, password string, urlStr string, body string) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", urlStr, bytes.NewBufferString(body))
	req.SetBasicAuth(username, password)
	return
}
