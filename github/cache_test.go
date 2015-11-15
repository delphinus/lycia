package github

import (
	"os"
	"path"
	"testing"
)

func withFakeCache(tmpRoot string) {
	dir, _ := path.Split(CachePath)
	_ = os.MkdirAll(dir, 0755)
	file, _ := os.OpenFile(CachePath, os.O_CREATE|os.O_WRONLY, 0666)

	defer func() { file.Close() }()

	file.WriteString(`[
		{
			"repository_url": "https://github.com/hoge/fuga",
			"branch": "hogehogeo",
			"pr_url": "https://github.com/hoge/fuga/pull/3"
		}
	]`)
}

func TestCacheLoad(t *testing.T) {
	withFakeEnv(t, func(tmpRoot string) {
		withFakeCache(tmpRoot)

		c := make(Cache)
		err := c.LoadCache()
		if err != nil {
			t.Errorf(`err found in LoadCache(): %s`, err)
		}

		if _, ok := c["https://github.com/hoge/fuga"]; !ok {
			t.Errorf(`cache for https://github.com/hoge/fuga is not found`)
		}

		if _, ok := c["https://github.com/hoge/fuga"]["hogehogeo"]; !ok {
			t.Errorf(`prUrl for hogehogeo is not found`)
		}
	})
}
