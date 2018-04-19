package bankservice

import (
	"context"

	"github.com/ivanmatyash/bank-golang/api"
)

func NewBankServer() api.BankServiceServer {
	return &bankServer{}
}

type bankServer struct {
}

func (s *bankServer) ListAccounts(ctx context.Context, req *api.RequestById) (*api.ResponseAccount, error) {
	return &api.ResponseAccount{[]*api.Account{&api.Account{1, 2, 3}}}, nil
}

func (s *bankServer) ReadAccount(ctx context.Context, req *api.RequestById) (*api.ResponseAccount, error) {
	return &api.ResponseAccount{[]*api.Account{&api.Account{req.GetId(), 2, 3}}}, nil
}

func (s *bankServer) ListClients(ctx context.Context, req *api.RequestById) (*api.ResponseClient, error) {
	return &api.ResponseClient{[]*api.Client{&api.Client{1, "ivan", "1", "11"}}}, nil
}

func (s *bankServer) ReadClient(ctx context.Context, req *api.RequestById) (*api.ResponseClient, error) {
	return &api.ResponseClient{[]*api.Client{&api.Client{1, "ivan", "1", "11"}}}, nil
}
