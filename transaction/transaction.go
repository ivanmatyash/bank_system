package transaction

import (
	"fmt"

	vld "github.com/ivanmatyash/bank-golang/validate"
)

type Transaction struct {
	id        int32
	diffMoney map[int32]int64 `db:"diff_money"`
	comment   string
	success   bool
}

func NewTransaction(cm string) (*Transaction, error) {
	tr := Transaction{id: 55, comment: cm}
	tr.diffMoney = make(map[int32]int64)
	if err := tr.validate(); err != nil {
		return nil, err
	}

	return &tr, nil
}

func (t *Transaction) Start() {
	fmt.Printf("Start transaction №%d.\n", t.id)
}

func (t *Transaction) End(ok bool) {
	t.success = ok
	fmt.Printf("End transcation №%d = %t.\n", t.id, t.success)
}

func (t *Transaction) DiffMoney() map[int32]int64 {
	return t.diffMoney
}
func (t *Transaction) AddDiffMoney(id int32, money int64) {
	t.diffMoney[id] = money
}

func (t Transaction) validate() error {
	err := vld.ValidateTransactionComment(t.comment)
	return err
}
