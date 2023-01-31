package errors

import "errors"

var (
	ErrOriginalURLExist = errors.New("originalURL exist")
	ErrNoSuchID         = errors.New("no such id")
	ErrURLDeleted       = errors.New("URL deleted")
)