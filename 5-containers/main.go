package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Up struct {
	Message string
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func up(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(Up{Message: "UP"})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(bytes))
}

func main() {
	var port = getenv("PORT", "8090")
	log.Printf("Server started on port %s", port)

	http.HandleFunc("/", up)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
