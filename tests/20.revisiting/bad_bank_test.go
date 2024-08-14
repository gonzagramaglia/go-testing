package revisiting

import (
	"math"
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertFloatEqual(t, newBalanceFor(riya), 200)
	AssertFloatEqual(t, newBalanceFor(chris), 0)
	AssertFloatEqual(t, newBalanceFor(adil), 175)
}

func AssertFloatEqual(t *testing.T, got, want float64) {
	t.Helper()
	const epsilon = 1e-9
	if math.Abs(got-want) > epsilon {
		t.Errorf("got %f, want %f", got, want)
	}
}
