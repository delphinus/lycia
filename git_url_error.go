package main

type GitUrlError string

func (err GitUrlError) Error() string {
	return string(err)
}
