package url_maker

import (
	"regexp"
)

var urlPattern = regexp.MustCompile(
	`^(?:(?P<scheme>https?|git|ssh)://)?` +
		`(?:(?P<username>[^@/]+)@)?` +
		`(?P<host>[^:/]+)[:/]` +
		`(?P<path>.*?)(?:\.git)?$`)

type MyError string

func (err MyError) Error() string {
	return string(err)
}

type GitUrl struct {
	RawUrl   string
	WebUrl   string
	Scheme   string
	Username string
	Host     string
	Path     string
}

func New(rawUrl string) (url GitUrl, err error) {
	url = GitUrl{RawUrl: rawUrl}
	err = url.Parse()
	return
}

func (url *GitUrl) Parse() (err error) {
	if !urlPattern.MatchString(url.RawUrl) {
		return MyError("this is not URL for git")
	}
	names := urlPattern.SubexpNames()[1:]
	m := urlPattern.FindStringSubmatch(url.RawUrl)[1:]
	matches := make(map[string]string)
	for i, str := range m {
		matches[names[i]] = str
	}
	url.Scheme = matches["scheme"]
	url.Username = matches["username"]
	url.Host = matches["host"]
	url.Path = matches["path"]
	return
}
