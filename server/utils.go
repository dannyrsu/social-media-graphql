package server

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
