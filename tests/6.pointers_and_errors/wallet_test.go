package wallet

import "testing"

func TestWallet(t *testing.T) {

	checkBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("depositing BTC", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdrawing BTC", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))

	})
}
