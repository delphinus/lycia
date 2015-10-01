package main

import (
	"github.com/delphinus35/lycia/git_url_maker"
	"log"
	"net/url"
	"os/exec"
)

func NewURL(dir string, ref string) (url *url.URL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		log.Fatalf("can not exec 'git remove -v' : %s", cmdErr)
	}
	maker, err := git_url_maker.New(string(out))
	url, err = url.Parse(maker.WebUrl)
	return
}
