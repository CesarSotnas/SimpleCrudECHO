package helpers

import "errors"

var (
	ErrMsgIdIsZero        = errors.New("id must be greater than 0")
	ErrMsgNotFound        = errors.New("not found")
	ErrMsgIdMustBeANumber = errors.New("id must be a number")
	ErrMsgEmptyValue      = errors.New("empty value")
)
