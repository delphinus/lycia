package main

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
)

// assignment for testing
var execCommand = exec.Command

func RemoteURL(dir string, ref string) (parsedURL *url.URL, err error) {
	cmd := execCommand("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		msg := fmt.Sprintf("can not exec 'git remove -v' : %s", cmdErr)
		log.Print(msg)
		err = MyError(msg)
	} else {
		maker, _ := UrlMaker(string(out))
		parsedURL, err = url.Parse(maker.WebUrl)
	}
	return
}
