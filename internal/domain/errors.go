package domain

import "errors"

// ErrPostNotFound is a custom error returned when a requested post cannot be found in the repository.
var ErrPostNotFound error

// init initializes the ErrPostNotFound variable with a descriptive error message.
func init() {
	ErrPostNotFound = errors.New("post not found")
}
