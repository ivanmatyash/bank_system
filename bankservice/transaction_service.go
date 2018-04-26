package bankservice

import (
	"log"
	"time"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

func (s *bankServer) StartTransaction(comment string) (*api.Transaction, error) {
	transaction := api.Transaction{}
	transaction.DiffMoney = make(map[int32]int64)
	transaction.Comment = comment
	if err := transaction.Validate(); err != nil {
		return nil, err
	}
	sqlstore.Mutex.Lock()
	log.Printf("Start transaction '%s'.\n", comment)

	return &transaction, nil
}

func (s *bankServer) EndTransaction(transaction *api.Transaction, ok bool) error {
	defer sqlstore.Mutex.Unlock()
	transaction.Success = ok
	var nextId int32
	err := sqlstore.Db.QueryRow("select nextval ('transactions_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		return err
	}
	transaction.Id = nextId
	transaction.Timestamp = time.Now().Unix()

	query := `
		INSERT INTO transactions(
			id,
			comment, 
			success,
			timestamp
		) VALUES(
			:id,
			:comment, 
			:success,
			:timestamp
		)`

	res, err := sqlstore.Db.NamedQuery(query, transaction)
	if err != nil {
		log.Println(err)
		res.Close()
		return err
	}
	res.Close()
	saveMoneyChangesInDB(transaction)

	log.Printf("End transaction '%s'.\n", transaction.Comment)

	return nil
}

func saveMoneyChangesInDB(t *api.Transaction) error {
	type record struct {
		TransactionId int32 `db:"transaction_id"`
		AccountId     int32 `db:"account_id"`
		Diff          int64 `db:"diff"`
	}

	records := make([]record, 0)

	for k, v := range t.DiffMoney {
		records = append(records, record{t.Id, int32(k), v})
	}
	query := `
		INSERT INTO money_changes(
			transaction_id,
			account_id, 
			diff
		) VALUES(
			:transaction_id,
			:account_id, 
			:diff
		)`
	for _, record := range records {
		res, err := sqlstore.Db.NamedQuery(query, record)
		if err != nil {
			log.Println(err)
			return err
		}
		res.Close()
	}
	log.Println("Accounts changed: ", t.DiffMoney)
	return nil
}
