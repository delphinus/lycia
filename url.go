package main

import (
	"log"
	"net/url"
	"os/exec"
	"regexp"
)

var gitUrlPattern = regexp.MustCompile("(?m)^origin\\s+(.*) \\(fetch\\)$")
var hasSchemePattern = regexp.MustCompile("^([^:]+)://")
var scpLikeUrlPattern = regexp.MustCompile("^([^@]+@)?([^:]+):/?(.+)$")

func NewURL(dir string, ref string) (url *url.URL, err error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		log.Fatalf("can not exec 'git remove -v' : %s", cmdErr)
	}
	outStr := string(out)
	url, err = url.Parse("https://github.com/powerline/powerline")
	return
}
