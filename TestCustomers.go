package main

import (
	"errors"
  "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

var collectors map[string]func (string) ([]Customer, error)

func init() {
	collectors = make(map[string]func (string) ([]Customer, error))

	collectors["default"] = func (fileName string) ([]Customer, error) {
		return nil, errors.New("No implementation found")
	}

	collectors["json"] = func (fileName string) ([]Customer, error) {
		customerFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}

		var customers []Customer

		err = json.Unmarshal(customerFile, &customers)
		if err != nil {
			return nil, err
		}

		return customers, nil
	}

	collectors["yml"] = func (fileName string) ([]Customer, error) {
		customerFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}

		var customers []Customer

		err = yaml.Unmarshal(customerFile, &customers)
		if err != nil {
			return nil, err
		}

		return customers, nil
	}
}

type Customer struct {
	AccountNumber float64
	Name string
	CreditLimit float64
	LocalCurrency string
	FavoriteStore string
}

func getCollectorFor(fileName string) func (string) ([]Customer, error) {
	splitStr := strings.Split(fileName, ".")
	fileExtension := splitStr[len(splitStr)-1]
	collector, ok := collectors[fileExtension]
	if ok {
		return collector
	}
	return collectors["default"]
}


func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		log.Fatalf("No file given\n")
	}

	collect := getCollectorFor(fileName)

	customers, err := collect(fileName)
	if err != nil {
		log.Fatalf("Encounted issue collecting customers from %s\nError: %s\n", fileName, err)
	}

	fmt.Printf("Using file:\t%s\n", fileName)

	fmt.Printf("How many customers:\t%d\n\n", len(customers))

	for _, customer := range(customers) {
		fmt.Println(customer)
	}
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

