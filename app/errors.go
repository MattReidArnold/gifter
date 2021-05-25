package app

import "errors"

var (
	ErrGroupNotFound        = errors.New("group not found")
	ErrGroupIDAlreadyExists = errors.New("group id already exists")
)
