package requests

import "errors"

var (
	ErrFieldsRequired      = errors.New("fields is required")
	ErrBadRequest          = errors.New("bad request")
	ErrorInternalServerErr = errors.New("internal server error")
	ErrUnauthorized        = errors.New("not authorized")
)
