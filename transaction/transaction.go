package transaction

import (
	"fmt"
	"log"

	"github.com/ivanmatyash/bank-golang/sqlstore"
	vld "github.com/ivanmatyash/bank-golang/validate"
)

type Transaction struct {
	Id        int32           `db:"id"`
	DiffMoney map[int32]int64 `db:"diff_money"`
	Comment   string          `db:"comment"`
	Success   bool            `db:"success"`
}

func NewTransaction(cm string) (*Transaction, error) {
	tr := Transaction{Comment: cm}
	tr.DiffMoney = make(map[int32]int64)
	if err := tr.validate(); err != nil {
		return nil, err
	}

	return &tr, nil
}

func (t *Transaction) Start() {
	fmt.Printf("Start transaction.\n")
}

func (t *Transaction) End(ok bool) {
	t.Success = ok
	var nextId int32
	err := sqlstore.Db.QueryRow("select nextval ('transactions_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
	}
	t.Id = nextId

	query := `
		INSERT INTO transactions(
			id,
			comment, 
			success
		) VALUES(
			:id,
			:comment, 
			:success
		)`

	_, err = sqlstore.Db.NamedQuery(query, t)
	t.saveMoneyChangesInDB()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("End transcation №%d = %t.\n", t.Id, t.Success)
}

func (t *Transaction) saveMoneyChangesInDB() error {
	type record struct {
		TransactionId int32 `db:"transaction_id"`
		AccountId     int32 `db:"account_id"`
		Diff          int64 `db:"diff"`
	}

	records := make([]record, 0)

	for k, v := range t.DiffMoney {
		records = append(records, record{t.Id, k, v})
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
		_, err := sqlstore.Db.NamedQuery(query, record)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (t *Transaction) AddDiffMoney(id int32, money int64) {
	t.DiffMoney[id] = money
}

func (t Transaction) validate() error {
	err := vld.ValidateTransactionComment(t.Comment)
	return err
}
