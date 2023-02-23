package schema

import (
	"encoding/json"
	"net/http"
)

func RenderJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	if body != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(body)
	} else {
		w.WriteHeader(statusCode)
	}
}
