package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/api"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("172.18.0.1:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewBankServiceClient(conn)

	accounts, err := c.ListAccountsByClient(ctx, &api.RequestById{1})
	if err != nil {
		log.Fatalln(err)
	}

	balance, err := c.GetBalance(ctx, &api.RequestAccount{Req: accounts.Result[0]})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(balance)
	_, err = c.ChangeBalance(ctx, &api.RequestAccountMoney{accounts.Result[0], 600})
	if err != nil {
		log.Fatalln(err)
	}
	balance, err = c.GetBalance(ctx, &api.RequestAccount{Req: accounts.Result[0]})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(balance)

	dst, _ := c.CreateAccount(ctx, &api.RequestAccount{Req: &api.Account{ClientId: 1}})

	a, err := c.TransferMoney(ctx, &api.RequestTransferMoney{accounts.Result[0], dst.Result[0], 10000000})
	fmt.Println(a)
}
