package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiHandler struct{}

var handler http.Handler

func main() {
	port := 8082

	http.HandleFunc("/", staticContentHandler)
	http.HandleFunc("/app/", dynamicContentHandler)

	handler = apiHandler{}
	http.Handle("/api/v.1/", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func staticContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Static content handler\n")
	//handler.ServeHTTP(w, r)
}

func dynamicContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dynamic content handler\n")
	//handler.ServeHTTP(w, r)
}
func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Api handler will be placed here\n")
}
