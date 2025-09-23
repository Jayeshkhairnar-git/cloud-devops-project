package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list/", http.StatusTemporaryRedirect)
}
