package main

import (
	"log"
	"net/url"
	"os/exec"
)

func RemoteURL(dir string, ref string) (url *url.URL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		log.Fatalf("can not exec 'git remove -v' : %s", cmdErr)
	}
	maker, err := UrlMaker(string(out))
	url, err = url.Parse(maker.WebUrl)
	return
}
