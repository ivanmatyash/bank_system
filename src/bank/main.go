package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/bankservice"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

var Addr string = "0.0.0.0:9091"

func main() {

	ln, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println(err)
	}

	err = sqlstore.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	api.RegisterBankServiceServer(server, bankservice.NewBankServer())

	server.Serve(ln)
}
