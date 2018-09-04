package localcache

import (
	"fmt"
)

const (
	NotSupportError = "not support"
	NotFoundError   = "not found"
)

type LocalCacheError struct {
	msg string
}

func NewLocalCacheError(msg string) error {
	return &LocalCacheError{
		msg: msg,
	}
}

func (lce *LocalCacheError) Error() string {
	return fmt.Sprintf("LocalCacheError %s", lce.msg)
}

func (lce *LocalCacheError) message() string {
	return lce.msg
}

func IsNoSupportError(e error) bool {
	lce, ok := e.(*LocalCacheError)
	if ok && lce.message() == NotSupportError {
		return true
	}
	return false
}

func IsFoundError(e error) bool {
	lce, ok := e.(*LocalCacheError)
	if ok && lce.message() == NotFoundError {
		return true
	}
	return false
}
