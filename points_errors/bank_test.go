package bank

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertInexistentError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("unexpected error")
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("expected error but got nil")
		}

		if err != want {
			t.Errorf("error result %s, want %s", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with enough balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertInexistentError(t, err)
	})

	t.Run("Withdraw without enough balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
