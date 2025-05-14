package views

import (
	"encoding/json"
	"net/http"
)

func (vw *Views) GetAPIInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status":"ok", "version": "v1.0.0"})
}
