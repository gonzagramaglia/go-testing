package wallet

import "fmt"

// It seems a bit overkill to create a struct for this.
// int is fine in terms of the way it works but it's not descriptive.
type Bitcoin int

type Stringer interface {
	String() string
} /*
This interface is defined in the fmt package
and lets you define how your type is printed when
used with the %s format string in prints.
*/
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
