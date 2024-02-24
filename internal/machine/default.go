package machine

import "cocom/internal/models"

type DefaultBehaviour struct {
	machine *machine
}

func (d *DefaultBehaviour) InsertMoney(_ int) error {
	return ErrTransactionNotValid
}

func (d *DefaultBehaviour) SelectSoda(_ string) error {
	return ErrTransactionNotValid
}

func (d *DefaultBehaviour) DispenseSoda(_ string) (*models.Soda, error) {
	return nil, ErrTransactionNotValid
}
