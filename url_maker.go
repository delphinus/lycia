package main

import (
	"fmt"
	"github.com/delphinus35/lycia/util"
	"regexp"
)

var urlPattern = regexp.MustCompile(
	`^(?:(?P<scheme>https?|git|ssh)://)?` +
		`(?:(?P<username>[^@/]+)@)?` +
		`(?P<host>[^:/]+)[:/]` +
		`(?P<path>.*?)(?:\.git)?$`)

type GitUrl struct {
	RawUrl   string
	WebUrl   string
	Scheme   string
	Username string
	Host     string
	Path     string
}

func UrlMaker(rawUrl string) (gitUrl GitUrl, err error) {
	gitUrl = GitUrl{RawUrl: rawUrl}
	err = gitUrl.Parse()
	if err == nil {
		gitUrl.makeWebUrl()
	}
	return
}

func (gitUrl *GitUrl) Parse() (err error) {
	if !urlPattern.MatchString(gitUrl.RawUrl) {
		return util.LyciaError("this is not URL for git")
	}
	names := urlPattern.SubexpNames()[1:]
	m := urlPattern.FindStringSubmatch(gitUrl.RawUrl)[1:]
	matches := make(map[string]string)
	for i, str := range m {
		matches[names[i]] = str
	}
	gitUrl.Scheme = matches["scheme"]
	gitUrl.Username = matches["username"]
	gitUrl.Host = matches["host"]
	gitUrl.Path = matches["path"]
	return
}

func (gitUrl *GitUrl) makeWebUrl() {
	gitUrl.WebUrl = fmt.Sprintf("https://%s/%s", gitUrl.Host, gitUrl.Path)
}
