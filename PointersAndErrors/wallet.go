package wallets

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
	// fmt.Printf("Depo wallet: %[1]p %#[1]v\n", w)
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if w.balance < value {
		return ErrInsufficientFunds
	}
	w.balance -= value
	return nil
}
