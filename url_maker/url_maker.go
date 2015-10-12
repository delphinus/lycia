package url_maker

import (
	"fmt"
	"regexp"
)

var urlPattern = regexp.MustCompile(
	`^(?:(?P<scheme>https?|git|ssh)://)?` +
		`(?:(?P<username>[^@/]+)@)?` +
		`(?P<host>[^:/]+)[:/]` +
		`(?P<path>.*?)(?:\.git)?$`)

type MyError string

func (self MyError) Error() (err string) {
	err = string(self)
	return
}

type GitUrl struct {
	RawUrl   string
	WebUrl   string
	Scheme   string
	Username string
	Host     string
	Path     string
}

func New(rawUrl string) (self GitUrl, err error) {
	self = GitUrl{RawUrl: rawUrl}
	err = self.Parse()
	if err != nil {
		self.makeWebUrl()
	}
	return
}

func (self *GitUrl) Parse() (err error) {
	if !urlPattern.MatchString(self.RawUrl) {
		return MyError("this is not URL for git")
	}
	names := urlPattern.SubexpNames()[1:]
	m := urlPattern.FindStringSubmatch(self.RawUrl)[1:]
	matches := make(map[string]string)
	for i, str := range m {
		matches[names[i]] = str
	}
	self.Scheme = matches["scheme"]
	self.Username = matches["username"]
	self.Host = matches["host"]
	self.Path = matches["path"]
	return
}

func (self *GitUrl) makeWebUrl() {
	self.WebUrl = fmt.Sprintf("%s://%s/%s")
}
