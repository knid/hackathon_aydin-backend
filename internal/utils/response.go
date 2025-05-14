package utils

import (
	"net/http"
	"encoding/json"
)

func JSONResponse(w http.ResponseWriter, status int, err any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

