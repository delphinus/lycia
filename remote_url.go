package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
)

var remoteUrlPattern = regexp.MustCompile(`(?m)^origin\s+(.*)\s+\(fetch\)$`)

type githubURL struct {
	Ref string
	*url.URL
}

func (g githubURL) SourceURL(path string, from int, to int) (sourceURL string) {
	sourceURL = g.String()
	if path != "" {
		sourceURL = fmt.Sprintf("%s/blob/%s/%s", sourceURL, g.Ref, path)
		if from != 0 {
			sourceURL = fmt.Sprintf("%s#L%d", sourceURL, from)
			if to != 0 {
				sourceURL = fmt.Sprintf("%s-L%d", sourceURL, to)
			}
		}
	} else if g.Ref != "master" {
		sourceURL = fmt.Sprintf("%s/tree/%s", g.Path, g.Ref)
	}
	return
}

func (g githubURL) IssueURL(num int) (issueURL string) {
	if num == 0 {
		issueURL = g.String() + "/issues"
	} else if num > 0 {
		issueURL = fmt.Sprintf("%s/issue/%d", g.String(), num)
	} else {
		issueURL = g.String()
	}
	return
}

func (g githubURL) PullrequestURL(num int) (pullrequestURL string) {
	if num == 0 {
		pullrequestURL = g.String() + "/pulls"
	} else if num > 0 {
		pullrequestURL = fmt.Sprintf("%s/pull/%d", g.String(), num)
	} else {
		pullrequestURL = g.String()
	}
	return
}

func RemoteURL(dir string, ref string) (parsed *githubURL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	outStr := string(out)

	if cmdErr != nil {
		msg := fmt.Sprintf("can not exec 'git remove -v' : %s", cmdErr)
		err = MyError(msg)

	} else if outStr == "" {
		err = MyError("git remote is not defined")

	} else if !remoteUrlPattern.MatchString(outStr) {
		msg := fmt.Sprintf("unknown git remote string: %s", outStr)
		err = MyError(msg)

	} else {
		rawUrl := remoteUrlPattern.FindStringSubmatch(outStr)[1]
		gitUrl, _ := UrlMaker(rawUrl)
		parsedURL, err := url.Parse(gitUrl.WebUrl)
		if err == nil {
			parsed = &githubURL{ref, parsedURL}
		}
	}
	return
}
