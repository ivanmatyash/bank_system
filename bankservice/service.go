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

func (s *bankServer) ReadClient(ctx context.Context, req *api.RequestById) (*api.ResponseClient, error) {
	client := api.Client{}
	err := sqlstore.Db.Get(&client, "SELECT * FROM clients WHERE id=$1", req.GetId())
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return &api.ResponseClient{[]*api.Client{&api.Client{}}}, status.Errorf(codes.NotFound, "Client %d not found.", req.GetId())
		}

		return nil, err
	}
	return &api.ResponseClient{[]*api.Client{&client}}, nil
}

func (s *bankServer) CreateClient(ctx context.Context, req *api.RequestClient) (*api.ResponseClient, error) {
	if err := req.Req.Validate(); err != nil {
		return &api.ResponseClient{[]*api.Client{&api.Client{}}}, err
	}
	var nextId int32
	err := sqlstore.Db.QueryRow("select nextval ('clients_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		return &api.ResponseClient{[]*api.Client{&api.Client{}}}, err
	}
	req.Req.Id = nextId

	query := `
		INSERT INTO clients(
			id,
			name, 
			email,
			phone
		) VALUES(
			:id,
			:name, 
			:email,
			:phone
		)`
	_, err = sqlstore.Db.NamedQuery(query, req.Req)

	if err != nil {
		return &api.ResponseClient{[]*api.Client{&api.Client{}}}, err
	}

	return &api.ResponseClient{[]*api.Client{req.Req}}, nil
}
