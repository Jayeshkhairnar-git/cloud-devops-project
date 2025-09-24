package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/items/", listItems).Methods("GET")
	r.HandleFunc("/items/", createItem).Methods("POST")
	r.HandleFunc("/items/{ID}", getItem).Methods("GET")
	r.HandleFunc("/items/{ID}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{ID}", deleteItem).Methods("DELETE")

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/item/", http.StatusTemporaryRedirect)
}
