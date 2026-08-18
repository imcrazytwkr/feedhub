package models

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Message    any
}

// Makes HTTPError compatible with the error interface
func (h *HTTPError) Error() string {
	return fmt.Sprintf("%d - %v", h.StatusCode, h.Message)
}

func NewHttpError(statusCode int, message any) error {
	h := &HTTPError{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
	}

	if message == nil {
		return h
	}

	httpError, ok := message.(HTTPError)
	if ok {
		h.Message = httpError.Message
		return h
	}

	h.Message = message
	return h
}
