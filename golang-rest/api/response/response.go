package response

import (
	"encoding/json"
	"net/http"
)

// messageResponse represents the structure of simple message responses.
type messageResponse struct {
	Message string `json:"message"`
}

// WriteOK writes a simple message response.
func WriteOK(w http.ResponseWriter, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(messageResponse{Message: msg})
}

// WriteJSON marshals data and writes a JSON response.
// On encoding failure, a 500 response is sent to the client.
func WriteJSON(w http.ResponseWriter, data any) error {
	buf, err := json.Marshal(data)
	if err != nil {
		// Ignore the returned error of the fallback error response.
		// The original marshal error is more useful for logging
		// since it's likely to be a bug.
		_ = WriteError(w, err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buf)
	return err
}
