package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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
	log.Printf("Requesting insult")
	resp, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en")
	if err != nil {
		// handle error
		log.Printf("Failed to request insult")
		log.Printf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Printf(string(body))
	bytes, _ := json.Marshal(Up{Message: "UP"})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(bytes))
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var port = getenv("PORT", "8090")
	log.Printf("Server started on port %s", port)

	http.HandleFunc("/", up)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
