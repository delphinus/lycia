package github

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func withFakeEnv(t *testing.T, block func(string)) {
	tmpRoot, err := ioutil.TempDir("", "lycia_")
	if err != nil {
		t.Fatalf("cloud not create tempdir: %s", err)
	}
	defer func() { os.RemoveAll(tmpRoot) }()

	ConfigPath = tmpRoot + "/.config/lycia/config.json"
	CachePath = tmpRoot + "/.config/lycia/cache.json"

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

func TestConfigLoad(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeConfig(tmpRoot)

		c := make(Config)
		err := c.Load()
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
