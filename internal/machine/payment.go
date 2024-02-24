package machine

type paymentState struct {
	DefaultBehaviour
	coins int
}

func NewPaymentState(machine *machine, coins int) State {
	return &paymentState{DefaultBehaviour{machine: machine}, coins}
}

func (p *paymentState) InsertMoney(coins int) error {
	p.coins += coins
	return nil
}

func (p *paymentState) SelectSoda(id string) error {
	soda, err := p.machine.findItem(id)
	if err != nil {
		return err
	}
	if soda.Quantity <= 0 {
		return ErrSodaRunningOut
	}
	if soda.Price > p.coins {
		return ErrNotEnoughMoney
	}
	soda.Quantity--
	p.machine.setState(NewDispenseState(p.machine, soda))
	return nil
}
