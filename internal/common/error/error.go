package error

import "errors"

var (
	ErrNotFound   = errors.New("resource not found")
	ErrInternal   = errors.New("internal server error")
	ErrBadRequest = errors.New("bad request")
)
