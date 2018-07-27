package web

import (
	"fmt"
	"net/http"
)

var mux *http.ServeMux
var customerGetSources map[string]func(http.ResponseWriter, *http.Request)

func init() {
	mux = http.NewServeMux()
	serviceConfig = &Config{}
	customerGetSources = map[string]func(http.ResponseWriter,*http.Request) {

	}
}

func Serve(config Config) (error) {
	*serviceConfig = config
	if len(serviceConfig.JsonRoute) == 2 {
		mux.HandleFunc(
			serviceConfig.DefaultRoute+"/"+serviceConfig.JsonRoute[0],
			customersHandler(serviceConfig.JsonRoute[1]),
		)
	}

	if len(serviceConfig.YamlRoute) == 2 {
		mux.HandleFunc(
			serviceConfig.DefaultRoute+"/"+serviceConfig.YamlRoute[0],
			customersHandler(serviceConfig.YamlRoute[1]),
		)
	}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil
}

func customersHandler(sourceFile string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			customersGetHandler(sourceFile)(w, r)
		default:
			invalidRequestHandler(w, r)
		}
	}
}

func invalidRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Invalid request")
}
