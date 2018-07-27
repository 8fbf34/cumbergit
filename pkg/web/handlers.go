package web

import (
	"fmt"
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	serviceConfig = &Config{}
}

func Serve(config Config) (error) {
	*serviceConfig = config
	mux.HandleFunc("/api/customers", customersHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		customersGetHandler(w, r)
	default:
		invalidRequestHandler(w, r)
	}
}

func invalidRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Invalid request")
}
