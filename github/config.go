package github

import (
	"encoding/json"
	"fmt"
	. "github.com/delphinus35/lycia/error"
	"io/ioutil"
	"os"
)

var ConfigPath = os.Getenv("HOME") + "/.config/lycia/config.json"

type Config map[string]SiteConfig

func (c Config) Load() (err error) {
	err = InitPath(ConfigPath)
	if err != nil {
		return
	}

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
		err = LyciaError(fmt.Sprintf("config path '%s' is corrupted: %s", ConfigPath, err))
		return
	}

	for _, siteConfig := range rawConfig {
		c[siteConfig.Host] = siteConfig
	}

	return
}

func (c Config) Save() (err error) {
	err = InitPath(ConfigPath)
	if err != nil {
		return
	}

	var rawConfig []SiteConfig
	for _, config := range c {
		rawConfig = append(rawConfig, config)
	}

	byt, err := json.Marshal(rawConfig)
	if err != nil {
		err = LyciaError(fmt.Sprintf("cannot encode config to JSON: %s", err))
		return
	}

	err = ioutil.WriteFile(ConfigPath, byt, 0644)
	if err != nil {
		err = LyciaError(fmt.Sprintf("cannot write config to file '%s': %s", ConfigPath, err))
	}
	return
}
