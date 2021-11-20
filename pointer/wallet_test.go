package wallet

import "testing"

// You can use errcheck to find missing checks
// go get -u github.com/kisielk/errcheck

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(5)
		assertBalance(t, wallet, want)
		assertNoError(t, err)

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialFunds := Bitcoin(15)
		wallet := Wallet{balance: initialFunds}
		err := wallet.Withdraw(Bitcoin(16))

		assertBalance(t, wallet, initialFunds)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		// Fatal stops the test
		t.Fatal("wanted an error but got none")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
