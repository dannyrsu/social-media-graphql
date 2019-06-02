package server

import (
	"encoding/json"
	"net/http"
)

func WriteJsonMessage(w http.ResponseWriter, message string) {
	response := map[string]interface{}{"message": message}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
