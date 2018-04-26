package bankservice

import (
	"context"

	"github.com/ivanmatyash/bank-golang/api"
)

func (s *bankServer) ChangeBalance(ctx context.Context, in *api.RequestAccountMoney) (*api.ResponseAccount, error) {
	transaction, err := s.StartTransaction("Balance changing.")
	if err != nil {
		return nil, err
	}
	in.Account.Balance += in.Money
	if err := in.Account.Validate(); err != nil {
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	res, err := s.UpdateAccount(ctx, &api.RequestAccount{in.Account, in.Account.Id})
	if err != nil {
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	transaction.DiffMoney[in.Account.Id] = in.GetMoney()

	if err := s.EndTransaction(transaction, true); err != nil {
		return nil, err
	}
	return &api.ResponseAccount{[]*api.Account{res.Result[0]}}, nil
}

func (s *bankServer) GetBalance(ctx context.Context, in *api.RequestAccount) (*api.ResponseMoney, error) {
	transaction, err := s.StartTransaction("Balance viewing.")
	if err != nil {
		return nil, err
	}
	res, err := s.ReadAccount(ctx, &api.RequestById{in.Req.Id})
	if err != nil {
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := s.EndTransaction(transaction, true); err != nil {
		return nil, err
	}
	return &api.ResponseMoney{res.Result[0].Balance}, nil
}
