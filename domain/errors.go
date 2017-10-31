package domain

import (
	"errors"
)

type Error struct {
	err error
	StatusCode int
}

func NewError(msg string, code int) *Error {
	e := Error{}
	e.err = errors.New(msg)
	e.StatusCode = code
	return &e
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Code() int {
	return e.StatusCode
}