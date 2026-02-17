package manager

import "errors"

var (
	ErrDuplicationToken = errors.New("duplicate token")
	ErrNotFound         = errors.New("bot not found")
)
