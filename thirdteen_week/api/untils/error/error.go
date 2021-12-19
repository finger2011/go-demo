package error

import (
	"github.com/pkg/errors"
)

const (
	ErrorCodeGrpc = 200000
)

type MyError interface {
	Wrap(msg string)
	Error() string
}

type ApiError struct {
	Code int
	Msg  string
	Err  error
}

func NewGRPCError(msg string) *ApiError {
	return &ApiError{
		Code: ErrorCodeGrpc,
		Msg:  msg,
		Err:  errors.New(msg),
	}
}

func NewError(msg string, code int) *ApiError {
	return &ApiError{
		Code: code,
		Msg:  msg,
		Err:  errors.New(msg),
	}
}

func (e *ApiError) Wrap(msg string) {
	e.Err = errors.Wrap(e.Err, msg)
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}
