package domain

import (
	"encoding/json"
	"errors"
	"micro-auth/serializer"
	"net/http"
)

type Error struct {
	err        error
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

func ErrToJSON(e error, code int) []byte {
	errResponse := serializer.Error{
		Status: http.StatusText(code),
		Error:  e.Error(),
	}

	errJSON, marshalErr := json.Marshal(errResponse)
	if marshalErr != nil {
		panic("code error: json marshalling failed for internally constructed struct")
	}

	return errJSON
}
