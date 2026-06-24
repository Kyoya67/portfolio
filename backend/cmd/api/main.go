package main

import (
	"log"
	"net/http"

	"github.com/kyoya67/portfolio/backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
