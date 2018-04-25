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
	accounts := []*api.Account{}
	err := sqlstore.Db.Select(&accounts, "SELECT * FROM accounts")
	if err != nil {
		return &api.ResponseAccount{[]*api.Account{}}, err
	}
	return &api.ResponseAccount{accounts}, nil
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

func (s *bankServer) CreateAccount(ctx context.Context, req *api.RequestAccount) (*api.ResponseAccount, error) {
	var nextId int32
	err := sqlstore.Db.QueryRow("select nextval ('accounts_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		return &api.ResponseAccount{[]*api.Account{&api.Account{}}}, err
	}
	req.Req.Id = nextId
	req.Req.Balance = 0

	query := `
		INSERT INTO accounts(
			id,
			client_id, 
			balance
		) VALUES(
			:id,
			:client_id, 
			:balance
		)`
	res, err := sqlstore.Db.NamedQuery(query, req.Req)
	if err != nil {
		return &api.ResponseAccount{[]*api.Account{&api.Account{}}}, err
	}
	res.Close()

	return &api.ResponseAccount{[]*api.Account{req.Req}}, nil
}

func (s *bankServer) UpdateAccount(ctx context.Context, req *api.RequestAccount) (*api.ResponseAccount, error) {
	if err := req.Req.Validate(); err != nil {
		return &api.ResponseAccount{[]*api.Account{&api.Account{}}}, err
	}

	req.GetReq().Id = req.GetId()
	query := `
		UPDATE accounts SET
			client_id = :client_id,
			balance = :balance
			WHERE id = :id
			`
	res, err := sqlstore.Db.NamedExec(query, req.GetReq())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Account %d not found.", req.GetId())
	}

	return &api.ResponseAccount{[]*api.Account{req.GetReq()}}, nil
}

func (s *bankServer) DeleteAccount(ctx context.Context, req *api.RequestById) (*api.ResponseAccount, error) {

	account := api.Account{}
	err := sqlstore.Db.Get(&account, "SELECT * FROM accounts WHERE id=$1", req.GetId())
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return &api.ResponseAccount{[]*api.Account{&api.Account{}}}, status.Errorf(codes.NotFound, "Account %d not found.", req.GetId())
		}

		return nil, err
	}

	_, err = sqlstore.Db.Exec("DELETE FROM accounts WHERE id=$1", req.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &api.ResponseAccount{[]*api.Account{&account}}, nil

}
