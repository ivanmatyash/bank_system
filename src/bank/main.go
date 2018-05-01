package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/bankservice"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

var Addr string = "bank.net:91"

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

func init() {
	if e := os.Getenv("BANK_SERVER_ADDR"); e != "" {
		Addr = e
	}

}
