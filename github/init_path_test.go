package github

import (
	"regexp"
	"testing"
)

func TestInitPathNotFound(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		ConfigPath = "/path/not/found"
		err := InitPath(ConfigPath)

		expected := "cannot mkdir"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestInitPath(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		err := InitPath(ConfigPath)

		if err != nil {
			t.Errorf(`got "%s" want "%s"`, err, "")
		}
	})
}
