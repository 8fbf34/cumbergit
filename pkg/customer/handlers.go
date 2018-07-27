package customer

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"errors"
)

type CollectorMap map[string]func(string) ([]Customer, error)
var collectors CollectorMap
var defaultKey = "default"

func init() {
	collectors = make(map[string]func(string)([]Customer, error))
	collectors["json"] = jsonCollector
	collectors["yml"] = yamlCollector
	collectors["mongo"] = mongoCollector
	collectors[defaultKey] = func(key string) ([]Customer, error) {
		return nil, errors.New("No implementation found")
	}
}

func (c CollectorMap) ContainsType(typeKey string) bool {
	_, ok := c[typeKey]
	return ok
}

func mongoCollector(url string) ([]Customer, error) {
	return nil, fmt.Errorf("Unimplemented")
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

func getCollectorFor(data string) func(string) ([]Customer, error) {
	var splitStr []string
	var dataType string

	// check if a URL
	splitStr = strings.Split(data, `://`)
	dataType = splitStr[0]
	if collector, ok := collectors[dataType]; ok {
		return collector
	}

	// check if a file
	splitStr = strings.Split(data, ".")
	if len(splitStr) > 1 {
		dataType = splitStr[len(splitStr)-1]
		log.Printf("File extension is:\t%s\n", dataType)
		if collector, ok := collectors[dataType]; ok {
			return collector
		}
	}

	return collectors[defaultKey]
}
