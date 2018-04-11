package client

import (
	"log"

	"github.com/ivanmatyash/bank-golang/sqlstore"
	"github.com/ivanmatyash/bank-golang/transaction"
)

type Client struct {
	Id    int32  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}

func NewClient(name string, email string, phone string) (cl *Client, err error) {
	var (
		tr *transaction.Transaction
	)
	if tr, err = transaction.NewTransaction("Create new client."); err != nil {
		return nil, err
	}

	tr.Start()
	cl = &Client{1, name, email, phone}

	var nextId int32
	err = sqlstore.Db.QueryRow("select nextval ('clients_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		tr.End(false)
		return nil, err
	}
	cl.Id = nextId

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
	_, err = sqlstore.Db.NamedQuery(query, cl)

	if err != nil {
		tr.End(false)
		return nil, err
	}

	tr.End(true)
	return cl, nil
}
