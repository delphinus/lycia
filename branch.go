package main

import (
	"github.com/delphinus35/lycia/util"
	"os/exec"
	"regexp"
)

var branchPattern = regexp.MustCompile(`(?m)^\* (.*)$`)

func DetectCurrentBranch(dir string) (branch string) {
	cmd := exec.Command("git", "branch")
	cmd.Dir = dir
	out, err := cmd.Output()

	if err != nil {
		err = util.LyciaError("can not exec 'git branch'")
		return
	}
	if !branchPattern.Match(out) {
		err = util.LyciaError("can not detect branch")
		return
	}
	branch = string(branchPattern.FindSubmatch(out)[1])
	return
}
