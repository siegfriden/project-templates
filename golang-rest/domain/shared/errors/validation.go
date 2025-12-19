package errors

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ErrValidation converts ozzo-validation errors into a domain error with field-specific metadata.
func ErrValidation(errs validation.Errors) Error {
	// Build metadata from validation errors.
	metadata := make(map[string]string, len(errs))
	for k, v := range errs {
		if err, ok := v.(validation.Error); ok {
			// Capitalize the first letter of the error messages and add period at the end.
			msg := err.Message()
			if len(msg) > 0 {
				msg = string(msg[0]-32) + msg[1:]
			}
			metadata[k] = msg + "."
		} else {
			// Some errors are not of type `validation.Error`, likely a code bug.
			// E.g. "cannot get the length of struct" when using length rules on a struct.
			// Return early with internal error.
			return Error{
				Kind:    KindInternalError,
				Message: DefaultInternalErrorMessage,
				Cause:   fmt.Errorf("validate field '%s': %w", k, v),
			}
		}
	}

	return Error{
		Kind:     KindInvalidInput,
		Message:  DefaultInvalidInputMessage,
		Cause:    errs,
		Metadata: metadata,
	}
}
