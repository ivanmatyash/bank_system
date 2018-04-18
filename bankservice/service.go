package bankservice

import (
	"context"

	"github.com/ivanmatyash/bank-golang/proto"
)

func NewBankServer() bank.BankServiceServer {
	return &bankServer{}
}

type bankServer struct {
}

func (s *bankServer) ListAccounts(ctx context.Context, req *bank.RequestById) (*bank.ResponseAccount, error) {
	return &bank.ResponseAccount{[]*bank.Account{&bank.Account{1, 2, 3}}}, nil
}

func (s *bankServer) ReadAccount(ctx context.Context, req *bank.RequestById) (*bank.ResponseAccount, error) {
	return &bank.ResponseAccount{[]*bank.Account{&bank.Account{req.GetId(), 2, 3}}}, nil
}
