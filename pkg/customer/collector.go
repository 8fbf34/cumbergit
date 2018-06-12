package customer

import (
	"fmt"
	"log"
)

func CollectAndProcess(fileName string) {
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

