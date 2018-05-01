package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/ivanmatyash/bank-golang/api"
)

const (
	baseURL = "/bank"
)

var (
	Addr     string = "gw_bank.net:80"
	BankAddr string = "bank.net:91"
)

func main() {

	handler, err := newBankHandler(context.Background(), BankAddr)
	if err != nil {
		log.Fatalln(err)
	}
	mux := http.NewServeMux()

	mux.Handle(baseURL+"/", http.StripPrefix(baseURL, handler))

	http.ListenAndServe(Addr, mux)
}

func newBankHandler(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) (http.Handler, error) {

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterBankServiceHandlerFromEndpoint(ctx, mux, addr, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func init() {
	if e := os.Getenv("GATEWAY_ADDR"); e != "" {
		Addr = e
	}
	if e := os.Getenv("BANK_SERVER_ADDR"); e != "" {
		BankAddr = e
	}
}
