package customer

import (
	"fmt"
)

type Customer struct {
	AccountNumber float64
	Name string
	CreditLimit float64
	LocalCurrency string
	FavoriteStore string
}

func (c Customer) String() string {
	s := ""
	s += fmt.Sprintf("Account Number:\t%f\n", c.AccountNumber)
	s += fmt.Sprintf("Name:\t%s\n", c.Name)
	s += fmt.Sprintf("Credit Limit:\t%f\n", c.CreditLimit)
	s += fmt.Sprintf("Local Currency:\t%s\n", c.LocalCurrency)
	s += fmt.Sprintf("Favorite Store:\t%s\n", c.FavoriteStore)
	return s
}

