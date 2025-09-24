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

type ItemHandler struct {
	items []Item
}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{
		items: []Item{},
	}
}

func (i *ItemHandler) listItems(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, i.items)
}

func (i *ItemHandler) createItem(w http.ResponseWriter, r *http.Request) {
	var request NewItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	item := Item{ID: uuid.NewString(), Name: request.Name, Description: request.Description}
	i.items = append(i.items, item)

	respondWithJSON(w, http.StatusCreated, item)
}

func (i *ItemHandler) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx := slices.IndexFunc(i.items, func(i Item) bool { return i.ID == vars["ID"] })
	if idx < 0 {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	respondWithJSON(w, http.StatusOK, i.items[idx])
}

func (i *ItemHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	var request NewItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	idx := slices.IndexFunc(i.items, func(i Item) bool { return i.ID == vars["ID"] })
	if idx < 0 {
		respondWithError(w, http.StatusNotFound, "")
		return
	}

	i.items[idx] = Item{
		ID:          vars["ID"],
		Name:        request.Name,
		Description: request.Description,
	}
}

func (i *ItemHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx := slices.IndexFunc(i.items, func(i Item) bool { return i.ID == vars["ID"] })
	if idx < 0 {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	i.items = slices.Delete(i.items, idx, 1)
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
