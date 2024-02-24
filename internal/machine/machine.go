package machine

import (
	"cocom/internal/models"
	"sync"
)

type machine struct {
	currState State
	sodas     []*models.Soda
	mState    sync.Mutex
}

func NewMachine(sodas []*models.Soda) VendingState {
	m := &machine{
		sodas: sodas,
	}
	m.currState = NewReadyState(m)
	return m
}

func (m *machine) InsertMoney(amount int) error {
	m.mState.Lock()
	defer m.mState.Unlock()
	return m.currState.InsertMoney(amount)
}

func (m *machine) SelectSoda(id string) error {
	m.mState.Lock()
	defer m.mState.Unlock()
	return m.currState.SelectSoda(id)
}

func (m *machine) DispenseSoda(id string) (*models.Soda, error) {
	m.mState.Lock()
	defer m.mState.Unlock()
	return m.currState.DispenseSoda(id)
}

func (m *machine) GetSodas() []*models.Soda {
	return m.sodas
}

func (m *machine) Restock(id string, amount int) error {
	soda, err := m.findItem(id)
	if err != nil {
		return err
	}
	soda.Quantity += amount
	return nil
}

func (m *machine) UpdatePrice(id string, newPrice int) error {
	soda, err := m.findItem(id)
	if err != nil {
		return err
	}
	soda.Price = newPrice
	return nil
}

func (m *machine) AddNewSoda(soda models.Soda) error {
	m.sodas = append(m.sodas, &soda)
	return nil
}

func (m *machine) findItem(id string) (*models.Soda, error) {
	for _, soda := range m.sodas {
		if soda.ID == id {
			return soda, nil
		}
	}
	return nil, ErrSodaNotFound
}

func (m *machine) setState(newState State) {
	m.currState = newState
}
