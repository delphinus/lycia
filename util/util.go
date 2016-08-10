package util

type LyciaError string

func (err LyciaError) Error() string {
	return string(err)
}
