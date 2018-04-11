package main

import (
	"fmt"
	"github.com/ivanmatyash/bank-golang/account"

	"github.com/ivanmatyash/bank-golang/client"
	//"github.com/ivanmatyash/bank-golang/account"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

func main() {
	err := sqlstore.InitDB()
	fmt.Println(err)
	client1, err := client.NewClient("Ivan", "test@gmail.com", "1010101")
	fmt.Println("ERROR=", err)
	acc1, err := account.NewAccount(client1.Id, 50)
	fmt.Println(acc1, err)
	err = acc1.Charge(20)
	acc1.Charge(30)
	acc1.Add(521)
	fmt.Println(err)
	//err = acc1.Add(35)
	//fmt.Println(acc1.Balance())
	//fmt.Println(err)
}
