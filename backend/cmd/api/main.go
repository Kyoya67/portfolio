package main

import (
	"log"
	"net/http"

	"github.com/kyoya67/portfolio/backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /works", handler.ListWorks)
	mux.HandleFunc("GET /works/{id}", handler.GetWork)
	mux.HandleFunc("GET /articles", handler.ListArticles)
	mux.HandleFunc("GET /articles/{id}", handler.GetArticle)
	mux.HandleFunc("GET /profile", handler.GetProfile)

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
