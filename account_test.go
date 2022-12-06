package account

import "testing"

func TestAccount(t *testing.T) {
	t.Run("returns 0 as initial cash account balance", func(t *testing.T) {
		account := Account{}

		expected := 0
		actual := account.ShowCashBalance()

		if actual != expected {
			t.Errorf("Expected account balance to be %d, but got %d", expected, actual)
		}
	})
	t.Run("deposits would not create error and increase balance", func(t *testing.T) {
		account := Account{}
		account.DepositCash(10)

		expected := 10
		actual := account.ShowCashBalance()

		if actual != expected {
			t.Errorf("Expected account balance to be %d, but got %d", expected, actual)
		}
	})
	t.Run("Withdraw an amount less than balance creates no error and decrease balance accordingly", func(t *testing.T) {
		account := Account{}
		account.DepositCash(10)
		account.WithdrawCash(7)

		expected := 3
		actual := account.ShowCashBalance()

		if actual != expected {
			t.Errorf("Expected account balance to be %d, but got %d", expected, actual)
		}
	})

}
