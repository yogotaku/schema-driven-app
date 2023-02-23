package schema

import (
	"encoding/json"
	"log"
	"net/http"
)

func RenderJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	log.Println(statusCode, body)
	if body != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(body)
	} else {
		w.WriteHeader(statusCode)
	}
}
