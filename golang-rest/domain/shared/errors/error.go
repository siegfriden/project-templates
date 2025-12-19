package errors

import "errors"

// Kind represents application error kinds.
// They are meant to be generic and map well to HTTP error codes.
type Kind string

const (
	KindInvalidInput           Kind = "INVALID_INPUT"           // 400
	KindAuthenticationRequired Kind = "AUTHENTICATION_REQUIRED" // 401
	KindPermissionDenied       Kind = "PERMISSION_DENIED"       // 403
	KindNotFound               Kind = "NOT_FOUND"               // 404
	KindConflict               Kind = "CONFLICT"                // 409
	KindInternalError          Kind = "INTERNAL_ERROR"          // 500
)

// Error is a composable error type with kind and optional metadata.
type Error struct {
	Kind     Kind
	Message  string // human-readable error message
	Cause    error
	Metadata map[string]string
}

// Error implements the error interface.
func (e Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return string(e.Kind)
}

// Unwrap returns the underlying cause of the error, if any.
func (e Error) Unwrap() error {
	return e.Cause
}

// New creates a new Error with the given kind and message.
func New(kind Kind, msg string) Error {
	return Error{Kind: kind, Message: msg}
}

// Wrap returns a new Error with the specified kind and message, wrapping the given error.
func Wrap(err error, kind Kind, msg string) Error {
	return Error{Kind: kind, Message: msg, Cause: err}
}

// AsError extracts the Error from an error.
func AsError(err error) Error {
	var e Error
	if errors.As(err, &e) {
		return e
	}
	// If the error is not of type Error, wrap it as an internal error.
	return Wrap(err, KindInternalError, DefaultInternalErrorMessage)
}
