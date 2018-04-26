package validate

import "fmt"

var (
	ErrorAccountNegativeBalance = "Error: Account balance cannot be negative!"
	ErrorEmptyString            = "string is empty."
	ErrorNegativeTransfer       = "Error: attempt to transfer negative amount of money (%d)."
)

func ValidateAccountBalance(balance int64) error {
	if balance < 0 {
		return fmt.Errorf("%s (%d)", ErrorAccountNegativeBalance, balance)
	}

	return nil
}

func ValidateLengthString(name string, s string, min int, max int) error {
	l := len(s)
	if l == 0 && min > 0 {
		return fmt.Errorf("%s %s", name, ErrorEmptyString)
	}
	return nil
}

func ValidateTransferMoney(money int64) error {
	if money <= 0 {
		return fmt.Errorf(ErrorNegativeTransfer, money)
	}
	return nil
}
