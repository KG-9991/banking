package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"pin-code" xml:"zip"`
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []customer{
		{Name: "kg", City: "NV", Zipcode: "1100"},
		{Name: "Rob", City: "ND", Zipcode: "2200"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}
