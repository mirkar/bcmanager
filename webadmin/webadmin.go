package main

import (
	"encoding/json"
	"fmt"
	"github.com/mirkar/bcmanager/prometheus"
	"log"
	"net/http"
)

type apiV1Handler struct{}
type apiV2Handler struct{}

var handlerApiV1 http.Handler
var handlerApiV2 http.Handler

func main() {
	port := 8082

	//reacthandler := http.FileServer(http.Dir("./static/react"))
	http.Handle("/reactdemo/", http.StripPrefix("/reactdemo/", http.FileServer(http.Dir("./static/react"))))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/app/", dynamicContentHandler)

	handlerApiV1 = apiV1Handler{}
	http.Handle("/api/v.1/", handlerApiV1)

	handlerApiV2 = apiV2Handler{}
	http.Handle("/api/v.2/", handlerApiV2)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func dynamicContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dynamic content handlerApiV1\n")
	//handlerApiV1.ServeHTTP(w, r)
}
func (h apiV1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := prometheus.Apache{
		Env:         "production",
		Datacenter:  "fairfield",
		Host:        "zlpv1830",
		Ingress:     7443,
		Acc:         "p1c1m292",
		Role:        "wca-cluster-ssl",
		Instance:    "130.4.170.16:9117",
		InstallPath: "/opt/app/apache/wca-cluster-ssl",
	}

	data, err := json.Marshal(response)

	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

func (h apiV2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := prometheus.Apache{
		Env:        "production",
		Datacenter: "alpharetta",
		Host:       "zlpv1733",
		Ingress:    6443,
		Acc:        "p1c1m260",
		Role:       "orphan-common-nossl",
		Instance:   "130.8.112.103:9117",
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
