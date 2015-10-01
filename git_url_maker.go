package git_url_maker

import (
	"regexp"
)

var gitUrlPattern = regexp.MustCompile(`^(?P<scheme>(?:https?|git|ssh)://)(?P<username>[^@]+@)?(?P<host>[^/]+)/(?P<path>.*?)(?:\.git)?`)

type MyError string

func (self MyError) Error() (err string) {
	err = string(self)
	return
}

type GitUrl struct {
	rawUrl   string
	scheme   string
	username string
	host     string
	path     string
}

func New(rawUrl string) (self GitUrl, err error) {
	self = GitUrl{rawUrl: rawUrl}
	err = self.Parse()
	return
}

func (self GitUrl) Parse() (err error) {
	if !gitUrlPattern.MatchString(self.rawUrl) {
		return MyError("this is not URL for git")
	}
	names := gitUrlPattern.SubexpNames()[1:]
	m := gitUrlPattern.FindStringSubmatch(self.rawUrl)
	matches := make(map[string]string)
	for i, str := range m {
		matches[names[i]] = str
	}
	self.scheme = matches["scheme"]
	self.username = matches["username"]
	self.host = matches["host"]
	self.path = matches["path"]
	return
}
