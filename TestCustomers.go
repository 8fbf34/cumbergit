package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type CustomerList struct {
	Customers []Customer
}

type Customer struct {
	AccountNumber int
	Name string
	CreditLimit int
	LocalCurrency string
	FavoriteStore string
}

var customerFileName string = "Customers.yml"

func main() {
	fmt.Println("Printing customer info")

	customerFile, err := ioutil.ReadFile(customerFileName)
	if err != nil {
		log.Fatalf("Encounted issue reading file %s", customerFileName)
	}

	customers := make(CustomerList)
	err = yaml.Unmarshal(customerFile, &customers)
	if err != nil {
		log.Fatalf("Encounted issue unmarshaling from %s", customerFileName)
	}

	for _, customer := range(customers.Customers) {
		fmt.Printf("Account Number:\t%d\n", customer.AccountNumber)
		fmt.Printf("Name:\t%s\n", customer.Name)
		fmt.Printf("Credit Limit:\t%d\n", customer.CreditLimit)
		fmt.Printf("Local Currency:\t%s\n", customer.LocalCurrency)
		fmt.Printf("Favorite Store:\t%s\n", customer.FavoriteStore)
	}
}

