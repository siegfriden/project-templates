package response

import (
	"encoding/json"
	"net/http"

	"gitlab.com/siegfriden/project-templates/golang-rest/domain/shared/errors"
)

// errorResponse represents the structure of error responses.
type errorResponse struct {
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// WriteError writes an error response with HTTP status derived from the domain error Kind.
// Non-domain errors default to 500 Internal Server Error.
func WriteError(w http.ResponseWriter, err error) error {
	domainErr := errors.AsError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus(domainErr.Kind))
	return json.NewEncoder(w).Encode(errorResponse{
		Message:  domainErr.Message,
		Metadata: domainErr.Metadata,
	})
}

// httpStatus maps domain error kinds to HTTP status codes.
func httpStatus(kind errors.Kind) int {
	switch kind {
	case errors.KindInvalidInput:
		return http.StatusBadRequest
	case errors.KindAuthenticationRequired:
		return http.StatusUnauthorized
	case errors.KindPermissionDenied:
		return http.StatusForbidden
	case errors.KindNotFound:
		return http.StatusNotFound
	case errors.KindConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
