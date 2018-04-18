package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ivanmatyash/bank-golang/proto"
	"google.golang.org/grpc"
)

const (
	baseURL = "/bank"
)

var (
	Addr     string = "0.0.0.0:8080"
	BankAddr string = "0.0.0.0:9091"
)

var handlers map[string]http.Handler

func initHandlers() {
	handlers = make(map[string]http.Handler)

	bankHandler, err := newBankHandler(context.Background(), BankAddr)
	if err != nil {
		log.Fatalln(err)
	}

	handlers["accounts"] = bankHandler
}

func httpServe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	mux := http.NewServeMux()
	url := strings.Split(r.URL.Path, "/")
	if len(url) < 3 {
		http.Error(w, "Error. Not implemented: "+r.URL.Path, 501)
		return
	}

	serviceName := url[2]

	if handler, ok := handlers[serviceName]; ok {
		mux.Handle(baseURL+"/", http.StripPrefix(baseURL, handler))
		mux.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Invalid address.", 501)
	return
}

func main() {
	initHandlers()
	mux := http.NewServeMux()
	mux.HandleFunc(baseURL+"/", httpServe)
	http.ListenAndServe(Addr, mux)
}

func newBankHandler(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) (http.Handler, error) {

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := bank.RegisterBankServiceHandlerFromEndpoint(ctx, mux, addr, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
