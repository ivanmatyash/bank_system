package bankservice

import (
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

func (s *bankServer) ListAccounts(ctx context.Context, req *api.RequestById) (*api.ResponseAccount, error) {
	return &api.ResponseAccount{[]*api.Account{&api.Account{1, 2, 3}}}, nil
}

func (s *bankServer) ReadAccount(ctx context.Context, req *api.RequestById) (*api.ResponseAccount, error) {
	account := api.Account{}
	err := sqlstore.Db.Get(&account, "SELECT * FROM accounts WHERE id=$1", req.GetId())
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return &api.ResponseAccount{[]*api.Account{&api.Account{}}}, status.Errorf(codes.NotFound, "Account %d not found.", req.GetId())
		}

		return nil, err
	}
	return &api.ResponseAccount{[]*api.Account{&account}}, nil
}
