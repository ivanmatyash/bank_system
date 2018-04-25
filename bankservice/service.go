package bankservice

import (
	"github.com/ivanmatyash/bank-golang/api"
)

func NewBankServer() api.BankServiceServer {
	return &bankServer{}
}

type bankServer struct {
}
