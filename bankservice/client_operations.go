package bankservice

import (
	"context"
	"github.com/ivanmatyash/bank-golang/validate"
	"log"

	"github.com/ivanmatyash/bank-golang/api"
)

func (s *bankServer) ChangeBalance(ctx context.Context, in *api.RequestAccountMoney) (*api.ResponseAccount, error) {
	transaction, err := s.StartTransaction("Balance changing.")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	in.Account.Balance += in.Money
	if err := in.Account.Validate(); err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	res, err := s.UpdateAccount(ctx, &api.RequestAccount{in.Account, in.Account.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	transaction.DiffMoney[in.Account.Id] = in.GetMoney()

	if err := s.EndTransaction(transaction, true); err != nil {
		log.Println(err)
		return nil, err
	}
	return &api.ResponseAccount{[]*api.Account{res.Result[0]}}, nil
}

func (s *bankServer) GetBalance(ctx context.Context, in *api.RequestAccount) (*api.ResponseMoney, error) {
	transaction, err := s.StartTransaction("Balance viewing.")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := s.ReadAccount(ctx, &api.RequestById{in.Req.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := s.EndTransaction(transaction, true); err != nil {
		log.Println(err)
		return nil, err
	}
	return &api.ResponseMoney{res.Result[0].Balance}, nil
}

func (s *bankServer) TransferMoney(ctx context.Context, in *api.RequestTransferMoney) (*api.ResponseAccount, error) {
	transaction, err := s.StartTransaction("Transfer money.")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := validate.ValidateTransferMoney(in.Money); err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	var (
		srcAccount *api.Account
		dstAccount *api.Account
	)

	res, err := s.ReadAccount(ctx, &api.RequestById{in.Src.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	srcAccount = res.Result[0]

	res, err = s.ReadAccount(ctx, &api.RequestById{in.Dst.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}
	dstAccount = res.Result[0]

	srcAccount.Balance -= in.Money
	if err := srcAccount.Validate(); err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	res, err = s.UpdateAccount(ctx, &api.RequestAccount{srcAccount, srcAccount.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	dstAccount.Balance += in.Money
	res, err = s.UpdateAccount(ctx, &api.RequestAccount{dstAccount, dstAccount.Id})
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			return nil, err
		}
		return nil, err
	}

	transaction.DiffMoney[srcAccount.Id] = -1 * in.Money
	transaction.DiffMoney[dstAccount.Id] = in.Money
	if err := s.EndTransaction(transaction, true); err != nil {
		return nil, err
	}

	return &api.ResponseAccount{[]*api.Account{srcAccount, dstAccount}}, nil

}
