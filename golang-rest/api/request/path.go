package request

import (
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/siegfriden/project-templates/golang-rest/domain/shared/errors"
)

const readIDErrMessage = "Invalid resource ID. Check the URL and try again."

// ReadID reads "{id}" path value from the URL and parses it into uuid.UUID.
// It works with `go-chi/chi` and standard `http` routers.
func ReadID(r *http.Request) (uuid.UUID, error) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, errors.KindInvalidInput, readIDErrMessage)
	}
	return id, nil
}
