package wallet

type Wallet struct {
	balance int
}

// https://gobyexample.com/pointers
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
