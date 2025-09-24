package handlers

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type NewItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var items = []Item{}

func listItems(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var request NewItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	item := Item{ID: uuid.NewString(), Name: request.Name, Description: request.Description}
	items = append(items, item)

	respondWithJSON(w, http.StatusCreated, item)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx := slices.IndexFunc(items, func(i Item) bool { return i.ID == vars["ID"] })
	if idx < 0 {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	respondWithJSON(w, http.StatusOK, items[idx])
}

func updateItem(w http.ResponseWriter, r *http.Request) {
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx := slices.IndexFunc(items, func(i Item) bool { return i.ID == vars["ID"] })
	if idx < 0 {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	items = slices.Delete(items, idx, 1)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
