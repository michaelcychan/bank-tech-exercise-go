package account

import "testing"

func TestAccount(t *testing.T) {
	t.Run("returns 0 as initial cash account balance", func(t *testing.T) {
		account := Account{}

		expected := Cash(0)
		actual := account.ShowCashBalance()

		assertCashBalance(t, expected, actual)
	})
	t.Run("deposits Cash would not create error and increase balance", func(t *testing.T) {
		account := Account{}
		account.DepositCash(10)

		expected := Cash(10)
		actual := account.ShowCashBalance()

		assertCashBalance(t, expected, actual)
	})
	t.Run("Withdraws a Cash amount less than balance creates no error and decrease balance accordingly", func(t *testing.T) {
		account := Account{}
		account.DepositCash(Cash(10))
		account.WithdrawCash(Cash(7))

		expected := Cash(2)
		actual := account.ShowCashBalance()

		assertCashBalance(t, expected, actual)
	})
}

func assertCashBalance(t testing.TB, expected, actual Cash) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected Cash account balance to be %s, but got %s", expected, actual)
	}
}
