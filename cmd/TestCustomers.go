package main

import (
	"log"
	"os"
	"github.com/8fbf34/cumbergit/pkg/customer"
)

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		log.Fatalf("No file given\n")
	}

	customer.CollectAndProcess(fileName)
}

