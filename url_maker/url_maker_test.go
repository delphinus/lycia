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
	var expected string
	m, _ := New("git://github.com/hoge/fuga.git")

	expected = "git"
	if m.Scheme != expected {
		t.Errorf(errFormat, m.Scheme, expected)
	}

	expected = ""
	if m.Username != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "github.com"
	if m.Host != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "hoge/fuga"
	if m.Path != expected {
		t.Errorf(errFormat, m.Username, expected)
	}
}
