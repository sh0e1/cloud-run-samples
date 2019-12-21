package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("httputil.DumpRequest: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Printf("dump: %s", string(dump))

		var m pubSubMessage
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		name := string(m.Message.Data)
		if name == "" {
			name = "World"
		}
		log.Printf("Hello %s!", name)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type pubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}
