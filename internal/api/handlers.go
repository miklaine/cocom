// Package api provides HTTP handlers for interacting with a vending machine.
// It includes operations such as listing sodas, inserting coins, purchasing sodas,
// and administrative functions like restocking and updating prices.
package api

import (
	"cocom/internal/machine"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Machine machine.VendingState
}

// NewHandler creates a new API handler with a reference to the inventory.
func NewHandler(vm machine.VendingState) *Handler {
	return &Handler{Machine: vm}
}

// RegisterRoutes registers the API routes.
func (h *Handler) RegisterRoutes() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/sodas", h.getSodas).Methods("GET")
	api.HandleFunc("/coins", h.insertCoins).Methods("POST")
	api.HandleFunc("/soda/{sodaID}/purchase", h.selectSoda).Methods("POST")

	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(authorize)

	admin.HandleFunc("/soda/{sodaID}/inventory", h.restock).Methods("POST")
	admin.HandleFunc("/soda/{sodaID}/price", h.updatePrice).Methods("POST")
	admin.HandleFunc("/sodas", h.addNewSoda).Methods("POST")

	http.Handle("/", r)
}

// authorize is a middleware function for securing admin routes.
func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Check for a specific token in the Authorization header
		token := r.Header.Get("Authorization")
		if token != "Bearer SuperSecretToken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// getSodas handles the retrieval of available sodas.
func (h *Handler) getSodas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Machine.GetSodas())
}

// insertCoins processes the insertion of coins into the vending machine.
func (h *Handler) insertCoins(w http.ResponseWriter, r *http.Request) {
	var insertMoneyRequest struct {
		Coins int `json:"coins"`
	}
	if err := json.NewDecoder(r.Body).Decode(&insertMoneyRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Machine.InsertMoney(insertMoneyRequest.Coins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// selectSoda handles the selection and purchase of a soda.
func (h *Handler) selectSoda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sodaID := vars["sodaID"]
	err := h.Machine.SelectSoda(sodaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	soda, err := h.Machine.DispenseSoda(sodaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(soda)
}
