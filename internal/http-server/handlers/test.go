package handlers

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string
	}{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
