package customer

import (
	"fmt"
)

type Customer struct {
	AccountNumber float64 `json:"accountNumber,omitempty",yaml:"accountNumber"`
	Name          string  `json:"name,omitempty",yaml:"name"`
	CreditLimit   float64 `json:"creditLimit,omitempty",yaml:"creditLimit"`
	LocalCurrency string  `json:"localCurrency,omitempty",yaml:"localCurrency"`
	FavoriteStore string  `json:"favoriteStore,omitempty",yaml:"favoriteStore"`
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
