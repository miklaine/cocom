package api

import (
	"cocom/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// restock handles the restocking of a specific soda by ID.
func (h *Handler) restock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sodaID := vars["sodaID"]
	var restockRequest struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&restockRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Machine.Restock(sodaID, restockRequest.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// updatePrice updates the price of a specific soda by ID.
func (h *Handler) updatePrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sodaID := vars["sodaID"]
	var updatePriceRequest struct {
		Price int `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updatePriceRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Machine.UpdatePrice(sodaID, updatePriceRequest.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// addNewSoda adds a new soda to the vending machine inventory.
func (h *Handler) addNewSoda(w http.ResponseWriter, r *http.Request) {
	var soda models.Soda
	if err := json.NewDecoder(r.Body).Decode(&soda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Machine.AddNewSoda(soda)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
