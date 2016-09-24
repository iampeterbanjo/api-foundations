package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)

	val := r.FormValue("name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Println("Starting server on port :80")
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
