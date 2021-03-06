package therrors

import (
	"strings"
)

// Error represents the error returned by server in response
type Error struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions"`
	Paths      []string               `json:"paths"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.Message
}

//NestErrorPaths is used to nest paths along with the error
func NestErrorPaths(e error, path string) error {
	err := ConvertError(e)

	newError := &Error{
		Paths:      []string{path},
		Extensions: err.Extensions,
		Message:    err.Message,
	}
	newError.Paths = append(newError.Paths, err.Paths...)

	return newError
}

type errorWithCode interface {
	Code() string
}

type errorWithExtensions interface {
	Extensions() map[string]interface{}
}

// ConvertError converts any error to jerrors.Error
func ConvertError(e error) *Error {
	err, ok := (e).(*Error)
	if !ok {
		code := "Unknown"
		if coder, ok := e.(errorWithCode); ok {
			code = coder.Code()
		}

		exts := make(map[string]interface{})

		if extender, ok := e.(errorWithExtensions); ok {
			exts = extender.Extensions()
		}

		exts["code"] = code

		return &Error{
			Paths:      []string{},
			Extensions: exts,
			Message:    e.Error(),
		}
	}

	return err
}

type MultiError struct {
	Errors []*Error
}

func (e *MultiError) Error() string {
	var s strings.Builder

	for _, e := range e.Errors {
		s.WriteString(e.Error())
		s.WriteString("\n")
	}

	return s.String()
}
