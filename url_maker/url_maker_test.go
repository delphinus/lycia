package url_maker

import (
	"testing"
)

var errFormat = `got "%s" want "%s"`

func TestParseError(t *testing.T) {
	_, err := New("some_bad_string")
	expected := "this is not URL for git"
	if err.Error() != expected {
		t.Errorf(errFormat, err.Error(), expected)
	}
}

func TestParseGit(t *testing.T) {
	m, _ := New("git://github.com/hoge/fuga.git")
	if m.Scheme != "git" {
		t.Errorf(errFormat, m.Scheme, "git")
	}
}
