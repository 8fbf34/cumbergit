package web

import (
	"github.com/8fbf34/cumbergit/pkg/customer"
	"net/http"
	"encoding/json"
)

func init() {
}

func customersGetHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := customer.Collect(serviceConfig.Source)

	if err != nil {
		json.NewEncoder(w).Encode(ServiceError{err.Error()})
	}

	json.NewEncoder(w).Encode(customers)
}
