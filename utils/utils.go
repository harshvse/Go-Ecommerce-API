package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if payload == nil {
		return fmt.Errorf("[-]Request has no payload")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(rw http.ResponseWriter, status int, v any) error {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	return json.NewEncoder(rw).Encode(v)
}

func WriteError(rw http.ResponseWriter, status int, err error) {
	WriteJSON(rw, status, map[string]string{"error": err.Error()})
}
