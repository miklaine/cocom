package machine

import "cocom/internal/models"

type State interface {
	InsertMoney(coins int) error
	SelectSoda(id string) error
	DispenseSoda(id string) (*models.Soda, error)
}

type VendingState interface {
	InsertMoney(coins int) error
	SelectSoda(id string) error
	DispenseSoda(id string) (*models.Soda, error)
	GetSodas() []*models.Soda
	Restock(id string, amount int) error
	UpdatePrice(id string, newPrice int) error
	AddNewSoda(soda models.Soda) error
}
