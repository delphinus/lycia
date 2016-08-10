package github

import (
	"encoding/json"
	"fmt"
	"github.com/delphinus35/lycia/util"
	"io/ioutil"
	"os"
)

var CachePath = os.Getenv("HOME") + "/.config/lycia/cache.json"

type BranchToUrl map[string]string
type Cache map[string]BranchToUrl

func (c Cache) LoadCache() (err error) {
	err = InitPath(CachePath)
	if err != nil {
		return
	}

	// If stat cannot be calculated, CachePath does not exist. This is not an error.
	if _, err = os.Stat(CachePath); err != nil {
		err = nil
		return
	}

	byt, err := ioutil.ReadFile(CachePath)
	if err != nil {
		err = util.LyciaError(fmt.Sprintf("access error to cache path '%s': %s", CachePath, err))
		return
	}

	var rawCache []PrUrlCache
	err = json.Unmarshal(byt, &rawCache)
	if err != nil {
		err = util.LyciaError(fmt.Sprintf("cache path '%s' is corrupted: %s", CachePath, err))
		return
	}

	for _, prUrlCache := range rawCache {
		var branchToUrl BranchToUrl
		if c[prUrlCache.RepositoryUrl] == nil {
			branchToUrl = make(BranchToUrl)
			c[prUrlCache.RepositoryUrl] = branchToUrl
		} else {
			branchToUrl = c[prUrlCache.RepositoryUrl]
		}
		branchToUrl[prUrlCache.Branch] = prUrlCache.PrUrl
	}

	return
}

func (c Cache) SaveCache() (err error) {
	err = InitPath(CachePath)
	if err != nil {
		return
	}

	var rawCache []PrUrlCache
	for repositoryUrl, cache := range c {
		for branch, prUrl := range cache {
			rawCache = append(rawCache, PrUrlCache{repositoryUrl, branch, prUrl})
		}
	}

	byt, err := json.Marshal(rawCache)
	if err != nil {
		err = util.LyciaError(fmt.Sprintf("cannot encode cache to JSON: %s", err))
		return
	}

	err = ioutil.WriteFile(CachePath, byt, 0644)
	if err != nil {
		err = util.LyciaError(fmt.Sprintf("cannot write cache to file '%s': %s", CachePath, err))
	}

	return
}
