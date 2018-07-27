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
	config.JsonRoute = []string{"json", "assets/Customers.json"}
	config.YamlRoute = []string{"yaml", "assets/Customers.yml"}

	err := web.Serve(config)
	if err != nil {
		log.Fatal(err)
		wg.Done()
	}
}