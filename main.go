package main

import "log"

type Account struct {
	FirstName string
	LastName  string
}

func main() {
	var accList []Account

	accList = append(accList, Account{
		FirstName: "Hiroyuki",
		LastName:  "Saito",
	})

	accList = append(accList, Account{
		FirstName: "Amane",
		LastName:  "Saito",
	})

	for i, name := range accList {
		log.Println(i, name)
	}

}
