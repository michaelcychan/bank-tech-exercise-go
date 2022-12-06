package account

import "fmt"

type Cash int

type Account struct {
	cashBalance Cash
}

func (a *Account) ShowCashBalance() Cash {
	return a.cashBalance
}

func (a *Account) DepositCash(cashAmount Cash) {
	a.cashBalance += cashAmount
}

func (a *Account) WithdrawCash(cashAmount Cash) {
	a.cashBalance -= cashAmount
}

func (c Cash) String() string {
	return fmt.Sprintf("$ %d", c)
}
