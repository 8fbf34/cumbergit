package main

import (
	"github.com/8fbf34/cumbergit/pkg/web"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	config := web.NewConfig()
	config.Source = "assets/Customers.json"

	err := web.Serve(config)
	if err != nil {
		log.Fatal(err)
		wg.Done()
	}
}