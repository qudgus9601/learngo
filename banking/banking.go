package banking

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string, balance int) *Account {
	account := Account{owner: owner, balance: balance}
	fmt.Println("Create Account Complete")
	return &account
}

/*
Dev
a : receiver라고 불린다.
보통 struct의 앞글자의 소문자로 쓴다.
*/
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

/*
Dev
a : receiver라고 불린다.
보통 struct의 앞글자의 소문자로 쓴다.
*/
func (a *Account) Balance() int{
	return a.balance
}

var errNoMoney error = errors.New("잔액이 부족합니다")

/*
Dev
a : receiver라고 불린다.
보통 struct의 앞글자의 소문자로 쓴다.
*/
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}