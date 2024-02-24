package machine

type readyState struct {
	DefaultBehaviour
}

func NewReadyState(machine *machine) State {
	return &readyState{
		DefaultBehaviour{machine: machine},
	}
}

func (r *readyState) InsertMoney(coins int) error {
	r.machine.setState(NewPaymentState(r.machine, coins))
	return nil
}
