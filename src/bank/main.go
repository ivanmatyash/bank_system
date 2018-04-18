package main

import (
	"fmt"
	"net"

	"github.com/ivanmatyash/bank-golang/proto"
	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/bankservice"
)

var Addr string = "0.0.0.0:9091"

func main() {

	ln, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()
	bank.RegisterBankServiceServer(server, bankservice.NewBankServer())

	server.Serve(ln)
}
