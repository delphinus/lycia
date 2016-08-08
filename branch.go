package main

import (
	"github.com/delphinus35/lycia/error"
	"os/exec"
	"regexp"
)

var branchPattern = regexp.MustCompile(`(?m)^\* (.*)$`)

func DetectCurrentBranch(dir string) (branch string) {
	cmd := exec.Command("git", "branch")
	cmd.Dir = dir
	out, err := cmd.Output()

	if err != nil {
		err = error.LyciaError("can not exec 'git branch'")
		return
	}
	if !branchPattern.Match(out) {
		err = error.LyciaError("can not detect branch")
		return
	}
	branch = string(branchPattern.FindSubmatch(out)[1])
	return
}
