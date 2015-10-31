package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
)

var remoteUrlPattern = regexp.MustCompile(`(?m)^origin\s+(.*)\s+\(fetch\)$`)

func RemoteURL(dir string, ref string) (parsedURL *url.URL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	outStr := string(out)

	if cmdErr != nil {
		msg := fmt.Sprintf("can not exec 'git remove -v' : %s", cmdErr)
		err = MyError(msg)

	} else if outStr == "" {
		err = MyError("git remote is not defined")

	} else if !remoteUrlPattern.MatchString(string(out)) {
		msg := fmt.Sprintf("unknown git remote string: %s", outStr)
		err = MyError(msg)

	} else {
		rawUrl := remoteUrlPattern.FindStringSubmatch(outStr)[1]
		gitUrl, _ := UrlMaker(rawUrl)
		parsedURL, err = url.Parse(gitUrl.WebUrl)
	}
	return
}
