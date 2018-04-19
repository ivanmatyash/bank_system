package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ivanmatyash/bank-golang/api"
	"google.golang.org/grpc"
)

const (
	baseURL = "/bank"
)

var (
	Addr     string = "0.0.0.0:8080"
	BankAddr string = "0.0.0.0:9091"
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
