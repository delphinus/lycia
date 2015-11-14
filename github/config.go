package github

import (
	"encoding/json"
	"fmt"
	. "github.com/delphinus35/lycia/error"
	"io/ioutil"
	"os"
	"path"
)

var ConfigPath = os.Getenv("HOME") + "/.config/lycia/config.json"

type Config map[string]SiteConfig

func (c Config) InitConfigPath() (err error) {
	dir, _ := path.Split(ConfigPath)
	stat, err := os.Stat(dir)
	if err != nil {
		err = LyciaError(fmt.Sprintf("access error to config path '%s': %s", ConfigPath, err))
		return
	}
	if !stat.IsDir() {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			err = LyciaError(fmt.Sprintf("cannot mkdir: '%s'", dir))
			return
		}
	}
	return
}

func (c Config) LoadConfig() (err error) {
	c.InitConfigPath()

	// If stat cannot be calculated, ConfigPath does not exist. This is not an error.
	if _, err = os.Stat(ConfigPath); err != nil {
		err = nil
		return
	}

	byt, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		err = LyciaError(fmt.Sprintf("access error to config path '%s': %s", ConfigPath, err))
		return
	}

	var rawConfig []SiteConfig
	err = json.Unmarshal(byt, &rawConfig)
	if err != nil {
		err = LyciaError(fmt.Sprintf("config path '%s' is corrupted"))
		return
	}

	for _, siteConfig := range rawConfig {
		c[siteConfig.Host] = siteConfig
	}

	return
}
