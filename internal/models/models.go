package models

// Soda represents virtual soda product
type Soda struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"cost"`
	Quantity    int    `json:"quantity"`
}
