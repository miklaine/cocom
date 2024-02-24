package machine

import (
	"cocom/internal/models"
)

type dispenseState struct {
	DefaultBehaviour
	soda *models.Soda
}

func NewDispenseState(machine *machine, soda *models.Soda) State {
	return &dispenseState{DefaultBehaviour{machine: machine}, soda}
}

func (d *dispenseState) DispenseSoda(id string) (*models.Soda, error) {
	if d.soda.ID != id {
		return nil, ErrNotExistForDispense
	}
	d.machine.setState(NewReadyState(d.machine))
	return d.soda, nil
}
