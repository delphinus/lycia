package main

import (
	"fmt"
	"github.com/delphinus35/lycia/error"
	"net/url"
	"os/exec"
	"regexp"
)

var remoteUrlPattern = regexp.MustCompile(`(?m)^origin\s+(.*)\s+\(fetch\)$`)

type remoteURL struct {
	Ref string
	*url.URL
}

func (r remoteURL) ToURL() *url.URL {
	return &url.URL{
		Scheme: r.Scheme,
		Host:   r.Host,
		Path:   r.Path,
	}
}

func (r remoteURL) SourceURL(ref string, path string, from int, to int) (sourceURL string) {
	r.Ref = ref
	sourceURL = r.String()
	if path != "" {
		sourceURL = fmt.Sprintf("%s/blob/%s/%s", sourceURL, r.Ref, path)
		if from != 0 {
			sourceURL = fmt.Sprintf("%s#L%d", sourceURL, from)
			if to != 0 {
				sourceURL = fmt.Sprintf("%s-L%d", sourceURL, to)
			}
		}
	} else if r.Ref != "master" {
		sourceURL = fmt.Sprintf("%s/tree/%s", r.Path, r.Ref)
	}
	return
}

func (r remoteURL) IssueURL(num int) (issueURL string) {
	if num == 0 {
		issueURL = r.String() + "/issues"
	} else if num > 0 {
		issueURL = fmt.Sprintf("%s/issues/%d", r.String(), num)
	} else {
		issueURL = r.String()
	}
	return
}

func RemoteURL(dir string) (parsed *remoteURL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	outStr := string(out)

	if cmdErr != nil {
		msg := fmt.Sprintf("can not exec 'git remove -v' : %s", cmdErr)
		err = error.LyciaError(msg)

	} else if outStr == "" {
		err = error.LyciaError("git remote is not defined")

	} else if !remoteUrlPattern.MatchString(outStr) {
		msg := fmt.Sprintf("unknown git remote string: %s", outStr)
		err = error.LyciaError(msg)

	} else {
		rawUrl := remoteUrlPattern.FindStringSubmatch(outStr)[1]
		gitUrl, _ := UrlMaker(rawUrl)
		parsedURL, err := url.Parse(gitUrl.WebUrl)
		if err == nil {
			parsed = &remoteURL{"", parsedURL}
		}
	}
	return
}
