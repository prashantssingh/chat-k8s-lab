package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	httpapi "go-chat/internal/http"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	h := httpapi.NewHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/readyz", h.Readyz)
	mux.HandleFunc("/send", h.Send)

	log.Println("go-chat listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
