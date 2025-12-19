package request

import (
	"encoding/json"
	"net/http"

	"gitlab.com/siegfriden/project-templates/golang-rest/domain/shared/errors"
)

const readJSONErrMessage = "Malformed or invalid JSON request body."

// ReadJSON decodes the JSON body into struct.
func ReadJSON[T any](r *http.Request) (*T, error) {
	dst := new(T)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // catch unexpected client fields (e.g., typos)
	if err := dec.Decode(dst); err != nil {
		return nil, errors.Wrap(err, errors.KindInvalidInput, readJSONErrMessage)
	}
	return dst, nil
}
