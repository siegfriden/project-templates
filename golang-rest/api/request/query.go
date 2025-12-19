package request

import (
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"gitlab.com/siegfriden/project-templates/golang-rest/domain/shared/errors"
)

const readURLQueryErrMessage = "Invalid URL query parameters."

// ReadURLQuery maps URL query into struct using `schema` tags.
// It supports primitive types, time.Time, and uuid.UUID.
func ReadURLQuery[T any](r *http.Request) (*T, error) {
	dst := new(T)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(false) // keep this false for consistent behavior with ReadJSON()
	err := dec.Decode(dst, r.URL.Query())
	if err == nil {
		return dst, nil
	}

	errs, ok := err.(schema.MultiError)
	if !ok {
		return nil, err // INTERNAL ERROR
	}

	// The MultiError map values aren't useful, so only the keys are returned.
	invalidKeys := make([]string, 0, len(errs))
	for k := range errs {
		invalidKeys = append(invalidKeys, k)
	}
	return nil, errors.Error{
		Kind:     errors.KindInvalidInput,
		Message:  readURLQueryErrMessage,
		Metadata: map[string]string{"invalid_keys": strings.Join(invalidKeys, ",")},
	}
}
