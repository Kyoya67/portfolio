package handler

import (
	"net/http"
	"strings"

	"github.com/kyoya67/portfolio/backend/internal/data"
)

func ListWorks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, data.Works)
}

func GetWork(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/works/")
	if id == "" || strings.Contains(id, "/") {
		writeNotFound(w)
		return
	}

	for _, work := range data.Works {
		if work.ID == id {
			writeJSON(w, http.StatusOK, work)
			return
		}
	}

	writeNotFound(w)
}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, data.Articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/articles/")
	if id == "" || strings.Contains(id, "/") {
		writeNotFound(w)
		return
	}

	for _, article := range data.Articles {
		if article.ID == id {
			writeJSON(w, http.StatusOK, article)
			return
		}
	}

	writeNotFound(w)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, data.Profile)
}
