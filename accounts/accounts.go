package accounts

import (
	"errors"
	"fmt"
)

// clean code!
var errnoMoney = errors.New("Can't withdraw")

//uppercase로 해야 import 를 할 수 있다

//Account struct
type Account struct {

	//lowercase : private
	//Uppercase : public
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {

	account := Account{owner: owner, balance: 0}
	return &account

}

// Method
// 처음에 구조체의 lowercase를 사용해줘야한다
// 복사해서 전달하지 말고 받은 Account를 그대로 사용
//Deposit x amount on your account
func (a *Account) Deposit(amount int) {

	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {

	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return errnoMoney
	}
	a.balance -= amount
	//error type 중 null 과 같다
	return nil

}
