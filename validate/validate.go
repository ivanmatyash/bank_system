package validate

import "fmt"

var (
	ErrorAccountIsNil           = "Error: Account is nil!"
	ErrorAccountCharge          = "Error: Insufficient funds on the account #%d(%d). Try to charge %d."
	ErrorAccountNegative        = "Error: money to charging/adding can't be negative! (%d)"
	ErrorAccountNegativeBalance = "Error: Account balance cannot be negative!"
	ErrorEmptyString            = "string is empty."
)

func ValidateAccountBalance(balance int64) error {
	if balance < 0 {
		return fmt.Errorf("%s (%d)", ErrorAccountNegativeBalance, balance)
	}

	return nil
}

func ValidateMoneyNegative(money int64) error {
	if money < 0 {
		return fmt.Errorf(ErrorAccountNegative, money)
	}
	return nil
}

func ValidateTransactionComment(comment string) error {
	err := validateLengthString("Comment to transaction", comment, 1, 255)
	return err
}

func validateLengthString(name string, s string, min int, max int) error {
	l := len(s)
	if l == 0 && min > 0 {
		return fmt.Errorf("%s %s", name, ErrorEmptyString)
	}
	return nil
}
