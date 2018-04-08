package main

import (
	"fmt"

	"github.com/ivanmatyash/bank-golang/account"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

func main() {
	err := sqlstore.InitDB()
	fmt.Println(err)
	acc1, err := account.NewAccount(50)
	fmt.Println(acc1, err)
	err = acc1.Charge(20)
	err = acc1.Add(35)
	fmt.Println(acc1.Balance())
	fmt.Println(err)
}
