package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if ok && name != "" {
		fmt.Fprintf(w, "Hello %s!", name)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Println("Starting server on port :80")

	m := mux.NewRouter()
	hey := m.PathPrefix("/hey").Subrouter()
	hey.HandleFunc("/{name}/", requestHandler)
	hey.HandleFunc("/{name}", requestHandler)
	hey.HandleFunc("/", requestHandler)

	http.Handle("/", m)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
