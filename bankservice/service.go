package bankservice

import (
	"context"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/sqlstore"
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
	clients := []*api.Client{}
	err := sqlstore.Db.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		return &api.ResponseClient{[]*api.Client{}}, err
	}
	return &api.ResponseClient{clients}, nil
}
