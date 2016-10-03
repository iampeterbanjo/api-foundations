package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"api-foundations/bootstrap"
)

// FirstName of Person
type FirstName struct {
	FirstName string `json:"firstname"`
}

// LastName of Person
type LastName struct {
	LastName string `json:"lastname"`
}

// FullName is FirstName and LastName of Person
type FullName struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getFirstName(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)
	value := r.FormValue("firstname")
	response := FirstName{FirstName: value}
	responseJSON := bootstrap.JSONEncode(response)
	fmt.Fprintf(w, responseJSON)
	fmt.Printf("[%.4f] Response with firstname: %s\n", bootstrap.Now(), value)
}

func getLastName(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)
	value := r.FormValue("lastname")
	response := LastName{LastName: value}
	responseJSON := bootstrap.JSONEncode(response)
	fmt.Fprintf(w, responseJSON)
	fmt.Printf("[%.4f] Response with lastname: %s\n", bootstrap.Now(), value)
}

func getFullName(w http.ResponseWriter, r *http.Request) {
	bootstrap.Reset()

	firstnameValue := r.FormValue("firstname")
	lastnameValue := r.FormValue("lastname")

	var firstname FirstName
	var lastname LastName
	var fullname FullName

	data := url.Values{}
	data.Add("firstname", firstnameValue)
	data.Add("lastname", lastnameValue)

	// fetch firstname
	fnChan := make(chan []byte, 1)
	go func() {
		fnURL := "http://localhost/firstname?" + data.Encode()
		fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), fnURL)
		fnResponse, _ := http.Get(fnURL)

		contents, _ := ioutil.ReadAll(fnResponse.Body)
		fnChan <- contents
	}()

	// fetch lastname
	lnChan := make(chan []byte, 1)
	go func() {
		lnURL := "http://localhost/fullname?" + data.Encode()
		fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), lnURL)
		lnResponse, _ := http.Get(lnURL)

		contents, _ := ioutil.ReadAll(lnResponse.Body)
		lnChan <- contents
	}()

	fnContents := <-fnChan
	_ = json.Unmarshal(fnContents, &firstname)
	fmt.Printf("%#v\n", &firstname)
	fullname.FirstName = firstname.FirstName

	lnContents := <-lnChan
	_ = json.Unmarshal(lnContents, &lastname)
	fullname.LastName = lastname.LastName

	fmt.Printf("[%.4f] Done fetching \n", bootstrap.Now())

	// return fullname response
	responseJSON := bootstrap.JSONEncode(fullname)
	fmt.Fprintf(w, responseJSON)
	fmt.Printf("[%.4f] Done with response: %#v\n", bootstrap.Now(), fullname)
}

func main() {
	fmt.Println("Starting server on port :9090")

	http.HandleFunc("/fullname", getFullName)
	http.HandleFunc("/firstname", getFirstName)
	http.HandleFunc("/lastname", getLastName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: %s", err)
	}
}
