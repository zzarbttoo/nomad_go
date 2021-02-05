package main

import (
	"fmt"
	"log"

	"github.com/zzarbttoo/learngo/accounts"
)

func main() {

	account := accounts.NewAccount("zzarbttoo")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)

	if err != nil {

		//프로그램 종료
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

}
