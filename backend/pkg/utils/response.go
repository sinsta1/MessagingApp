package utils

import (
    "encoding/json"
    "net/http"
)

// SendJSONResponse sends a JSON response to the client
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

// SendErrorResponse sends an error response to the client
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    SendJSONResponse(w, statusCode, map[string]string{"error": message})
}
