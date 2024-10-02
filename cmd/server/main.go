package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)

	log.Println(fmt.Sprintf("Server is running on : http://localhost:%s", PORT))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Tic Tac Toe 🧙🏽"}`))
}
