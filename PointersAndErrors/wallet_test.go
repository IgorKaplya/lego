package wallets

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	var wallet = Wallet{balance: 7}
	// fmt.Printf("Test wallet: %[1]p %#[1]v\n", &wallet)

	wallet.Deposit(3)

	assertBalance(t, &wallet, 10)
}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 7}
		err := wallet.Withdraw(3)
		assertHasNoError(t, err)
		assertBalance(t, &wallet, 4)
	})

	t.Run("Overdraft", func(t *testing.T) {
		var wallet = Wallet{balance: 5}
		var err = wallet.Withdraw(3)
		assertErrorIs(t, err, ErrInsufficientFunds)
		assertBalance(t, &wallet, 2)
	})
}

func assertHasNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("err %q", err)
	}
}

func assertErrorIs(t *testing.T, err error, want error) {
	t.Helper()
	if errors.Is(err, want) {
		t.Errorf("err %q want %q", err, want)
	}
}

func assertBalance(t testing.TB, wallet *Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
