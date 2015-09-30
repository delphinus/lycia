package main

import (
	"net/url"
)

func NewURL(dir string) (url *url.URL, err error) {
	url, err = url.Parse("https://github.com/powerline/powerline")
	return
}
