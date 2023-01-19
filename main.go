package main

import "log"

type Account struct {
	FirstName string
	LastName  string
	Balance   int
}

type Balance interface {
	GetBalance() int
	IncreaseBalance()
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) IncreaseBalance() {
	a.Balance = a.Balance + 50
	log.Println(a.Balance)
}

func main() {
	var accList []Account

	accList = append(accList, Account{
		FirstName: "Hiroyuki",
		LastName:  "Saito",
		Balance:   0,
	})

	accList = append(accList, Account{
		FirstName: "Amane",
		LastName:  "Saito",
		Balance:   100,
	})

	for _, acc := range accList {
		acc.IncreaseBalance()
	}

	for i, acc := range accList {
		log.Println(i, acc.FirstName, acc.GetBalance())
	}

}
