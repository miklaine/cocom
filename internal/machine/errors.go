package machine

import "errors"

var (
	ErrTransactionNotValid = errors.New("transaction not valid")
	ErrSodaNotFound        = errors.New("soda not found")
	ErrSodaRunningOut      = errors.New("soda running out")
	ErrNotEnoughMoney      = errors.New("not enough money")
	ErrNotExistForDispense = errors.New("not exist for dispense")
)
