package account

type Account struct {
	cashBalance int
}

func (a *Account) ShowCashBalance() int {
	return a.cashBalance
}

func (a *Account) DepositCash(cashAmount int) {
	a.cashBalance += cashAmount
}

func (a *Account) WithdrawCash(cashAmount int) {
	a.cashBalance -= cashAmount
}
