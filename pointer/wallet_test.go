package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(1)

	got := wallet.Balance()
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
