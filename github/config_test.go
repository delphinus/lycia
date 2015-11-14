package github

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"testing"
)

func withFakeEnv(t *testing.T, block func(string)) {
	tmpRoot, err := ioutil.TempDir("", "lycia_")
	if err != nil {
		t.Fatalf("cloud not create tempdir: %s", err)
	}
	defer func() { os.RemoveAll(tmpRoot) }()

	ConfigPath = tmpRoot + "/.config/lycia/config.json"

	block(tmpRoot)
}

func withFakeConfig(tmpRoot string) {
	dir, _ := path.Split(ConfigPath)
	_ = os.MkdirAll(dir, 0755)
	file, _ := os.OpenFile(ConfigPath, os.O_CREATE|os.O_WRONLY, 0666)

	defer func() { file.Close() }()

	file.WriteString(`[
		{
			"host": "github.com",
			"access_token": "hogehogeo"
		}
	]`)
}

func TestInitConfigPathNotFound(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		ConfigPath = "/path/not/found"
		c := make(Config)
		err := c.InitConfigPath()

		expected := "access error to config path"
		if ok, _ := regexp.MatchString(expected, err.Error()); !ok {
			t.Errorf(`got "%s" want "%s"`, err, expected)
		}
	})
}

func TestInitConfigPath(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeConfig(tmpRoot)

		c := make(Config)
		err := c.LoadConfig()
		if err != nil {
			t.Errorf(`err found in LoadConfig(): %s`, err)
		}

		if _, ok := c["github.com"]; !ok {
			t.Errorf(`setting for github.com is not found`)
		}

		if c["github.com"].AccessToken != "hogehogeo" {
			t.Errorf(`access_token for github.com is not found`)
		}
	})
}
