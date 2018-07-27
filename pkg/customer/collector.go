package customer

import (
	"fmt"
	"log"
)

func CollectAndProcess(fileName string) {
	collected, err := Collect(fileName)
	if err != nil {
		log.Printf("Issue encountered:\t%s\n", err)
		return
	}
	Process(collected, fileName)
}

func Collect(fileName string) ([]Customer, error) {
	collector := getCollectorFor(fileName)
	return collector(fileName)
}

func Process(customers []Customer, fileName string) {
	fmt.Printf("Using file:\t%s\n", fileName)

	fmt.Printf("How many customers:\t%d\n\n", len(customers))

	for _, customer := range customers {
		fmt.Println(customer)
	}
}
