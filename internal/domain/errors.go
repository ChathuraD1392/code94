package domain

import "errors"

var ErrPostNotFound error

func init() {
	ErrPostNotFound = errors.New("post not found")
}
