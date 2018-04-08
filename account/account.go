package account

import (
	"fmt"

	"github.com/ivanmatyash/bank-golang/transaction"
	vld "github.com/ivanmatyash/bank-golang/validate"
)

type Account struct {
	id      int32
	balance int64
}

func NewAccount(balance int64) (acc *Account, err error) {
	var (
		tr *transaction.Transaction
	)
	if tr, err = transaction.NewTransaction("Create new account."); err != nil {
		return nil, err
	}

	tr.Start()
	acc = &Account{1, balance}
	err = acc.validate()
	if err = acc.validate(); err != nil {
		tr.End(false)
		return nil, err
	}
	tr.End(true)
	return acc, nil
}

func (a *Account) Balance() int64 {
	return a.balance
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
	startBalance := a.balance
	a.balance -= money
	tr.AddDiffMoney(a.id, a.balance-startBalance)
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
	startBalance := a.balance
	a.balance += money
	tr.AddDiffMoney(a.id, a.balance-startBalance)
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
	startBalance := a.balance
	a.balance += money
	tr.AddDiffMoney(a.id, a.balance-startBalance)
	tr.End(true)
	return err
}

func (a *Account) validate() error {
	return vld.ValidateAccountBalance(a.balance)
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
	if a.balance < money {
		return fmt.Errorf(vld.ErrorAccountCharge, a.id, a.balance, money)
	}
	return nil
}
