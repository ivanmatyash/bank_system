package account

import (
	"fmt"
	"log"

	"github.com/ivanmatyash/bank-golang/sqlstore"
	"github.com/ivanmatyash/bank-golang/transaction"
	vld "github.com/ivanmatyash/bank-golang/validate"
)

type Account struct {
	Id       int32 `db:"id"`
	ClientId int32 `db:"client_id"`
	Balance  int64 `db:"balance"`
}

func NewAccount(clientId int32, balance int64) (acc *Account, err error) {
	var (
		tr *transaction.Transaction
	)
	if tr, err = transaction.NewTransaction("Create new account."); err != nil {
		return nil, err
	}

	tr.Start()
	acc = &Account{1, clientId, balance}
	if err = acc.validate(); err != nil {
		tr.End(false)
		return nil, err
	}
	var nextId int32
	err = sqlstore.Db.QueryRow("select nextval ('accounts_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		tr.End(false)
		return nil, err
	}
	acc.Id = nextId

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

	_, err = sqlstore.Db.NamedQuery(query, acc)

	if err != nil {
		log.Println(err)
		tr.End(false)
		return nil, err
	}

	tr.End(true)
	return acc, nil
}

func (a *Account) Charge(money int64) (err error) {
	var (
		tr *transaction.Transaction
	)

	if tr, err = transaction.NewTransaction("Charge money from account."); err != nil {
		return err
	}
	tr.Start()
	if err = a.validateChargeMoney(money); err != nil {
		tr.End(false)
		return err
	}
	startBalance := a.Balance
	a.Balance -= money

	query := `
		UPDATE accounts SET
			balance = :balance 
			WHERE id=:id
			`
	_, err = sqlstore.Db.NamedExec(query, a)
	if err != nil {
		log.Println(err)
		tr.End(false)
		return err
	}

	tr.AddDiffMoney(a.Id, a.Balance-startBalance)
	tr.End(true)
	return err
}

func (a *Account) Add(money int64) (err error) {
	var (
		tr *transaction.Transaction
	)

	if tr, err = transaction.NewTransaction("Add money to account."); err != nil {
		return err
	}
	tr.Start()
	if err = a.validateMoneyToOperations(money); err != nil {
		tr.End(false)
		return err
	}
	startBalance := a.Balance
	a.Balance += money

	query := `
		UPDATE accounts SET
			balance = :balance 
			WHERE id=:id
			`
	_, err = sqlstore.Db.NamedExec(query, a)
	if err != nil {
		log.Println(err)
		tr.End(false)
		return err
	}

	tr.AddDiffMoney(a.Id, a.Balance-startBalance)
	tr.End(true)
	return err
}

func (a *Account) Transfer(idAccount int32, money int64) (err error) {
	var (
		tr *transaction.Transaction
	)

	if tr, err = transaction.NewTransaction("Transfer money from account to another account."); err != nil {
		return err
	}
	tr.Start()
	if err = a.validateChargeMoney(money); err != nil {
		tr.End(false)
		return err
	}
	// TODO: Add check of existing account with idAccount
	// TODO: Add creating of object
	startBalance := a.Balance
	a.Balance += money
	tr.AddDiffMoney(a.Id, a.Balance-startBalance)
	tr.End(true)
	return err
}

func (a *Account) validate() error {
	return vld.ValidateAccountBalance(a.Balance)
}

func (a *Account) validateMoneyToOperations(money int64) error {
	if a == nil {
		return fmt.Errorf(vld.ErrorAccountIsNil)
	}
	if err := vld.ValidateMoneyNegative(money); err != nil {
		return err
	}
	return nil
}

func (a *Account) validateChargeMoney(money int64) error {
	if err := a.validateMoneyToOperations(money); err != nil {
		return err
	}
	if a.Balance < money {
		return fmt.Errorf(vld.ErrorAccountCharge, a.Id, a.Balance, money)
	}
	return nil
}
