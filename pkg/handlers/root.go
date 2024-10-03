package handlers

import "net/http"

func HandleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "This is tttm-rango by Domen Perko"}`))
}
