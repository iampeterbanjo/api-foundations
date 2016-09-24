package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)

	val := r.URL.Query().Get(":name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Println("Starting server on port :80")

	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(requestHandler))
	m.Get("/", http.HandlerFunc(requestHandler))

	http.Handle("/", m)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
