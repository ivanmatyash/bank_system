package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/bankservice"
)

var Addr string = "bank.net:91"

func main() {

	ln, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println(err)
	}

	//err = sqlstore.InitDB()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	server := grpc.NewServer()
	api.RegisterBankServiceServer(server, bankservice.NewBankServer())

	server.Serve(ln)
}
