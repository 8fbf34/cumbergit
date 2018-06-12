package customer

import (
	"errors"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
	"gopkg.in/yaml.v2"
)

var collectors map[string]func (string) ([]Customer, error)

func init() {
	collectors = make(map[string]func (string) ([]Customer, error))

	collectors["default"] = func (fileName string) ([]Customer, error) {
		return nil, errors.New("No implementation found")
	}

	collectors["json"] = jsonCollector

	collectors["yml"] = yamlCollector
}

func jsonCollector(fileName string) ([]Customer, error) {
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

func yamlCollector(fileName string) ([]Customer, error) {
	customerFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var customers []Customer

	err = yaml.Unmarshal(customerFile, &customers)
	if err != nil {
		return nil, err
	}

	log.Printf("First customer is:\t%s\n", customers[0])

	return customers, nil
}

func getCollectorFor(fileName string) func (string) ([]Customer, error) {
	splitStr := strings.Split(fileName, ".")
	fileExtension := splitStr[len(splitStr)-1]
	log.Printf("File extension is:\t%s\n", fileExtension)
	collector, ok := collectors[fileExtension]
	if ok {
		return collector
	}
	return collectors["default"]
}

