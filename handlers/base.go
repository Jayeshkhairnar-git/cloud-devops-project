package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
)

func New() *mux.Router {
	r := mux.NewRouter()

	s := stores.NewMemoryItemStore()

	c := NewItemHandler(s)

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/items/", c.listItems).Methods("GET")
	r.HandleFunc("/items/", c.createItem).Methods("POST")
	r.HandleFunc("/items/{ID}", c.getItem).Methods("GET")
	r.HandleFunc("/items/{ID}", c.updateItem).Methods("PUT")
	r.HandleFunc("/items/{ID}", c.deleteItem).Methods("DELETE")

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/item/", http.StatusTemporaryRedirect)
}
