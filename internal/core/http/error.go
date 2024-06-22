package httpcore

import (
	"reflect"

	"github.com/pkg/errors"
)

var (
	ErrBadStatusCode = errors.New("bad status code")
)

type response interface {
	StatusCode() int
	Body() []byte
	IsError() bool
}

func GetHandleErrorFunc[T any](api, method string, defaultValue T) func(err error, response response) (T, error) {
	return func(err error, response response) (T, error) {
		if response == nil || reflect.ValueOf(response).IsNil() {
			return defaultValue, NewHTTPError(err, api, method, 0, "")
		}
		return defaultValue, NewHTTPError(err, api, method, response.StatusCode(), string(response.Body()))
	}
}

func GetHandleErrorFuncNoRet(api, method string) func(err error, response response) error {
	return func(err error, response response) error {
		if response == nil {
			return NewHTTPError(err, api, method, 0, "")
		}
		return NewHTTPError(err, api, method, response.StatusCode(), string(response.Body()))
	}
}

func HandleHTTPError(err error, response response) error {
	if err != nil {
		return err
	}

	if response.IsError() {
		return ErrBadStatusCode
	}

	return nil
}

type HTTPError struct {
	API        string
	Method     string
	StatusCode int
	Message    string
	err        error
}

func NewHTTPError(err error, api, method string, statusCode int, message string) *HTTPError {
	return &HTTPError{
		API:        api,
		Method:     method,
		StatusCode: statusCode,
		Message:    message,
		err:        err,
	}
}

func (e *HTTPError) Error() string {
	return errors.Wrapf(
		e.err,
		"api %s error: %s: status code: %d, massage: %s",
		e.API,
		e.Method,
		e.StatusCode,
		e.Message,
	).Error()
}
