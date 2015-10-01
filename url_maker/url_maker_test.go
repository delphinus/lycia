package url_maker

import (
	"testing"
)

func TestParseError(t *testing.T) {
	_, err := New("some_bad_string")
	expected := "this is not URL for git"
	if err.Error() != expected {
		t.Errorf("got \"%s\" want \"%s\"", err.Error(), expected)
	}
}
