package api

import (
	"log"

	"github.com/ivanmatyash/bank-golang/validate"
)

func (c *Client) Validate() error {
	clientData := make(map[string]string)
	clientData["Client Name"] = c.Name
	clientData["Client email"] = c.Email
	clientData["Client phone"] = c.Phone
	for name, value := range clientData {
		if err := validate.ValidateLengthString(name, value, 1, 255); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (a *Account) Validate() error {
	if err := validate.ValidateAccountBalance(a.Balance); err != nil {
		return err
	}
	return nil
}
