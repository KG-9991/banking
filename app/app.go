package main

import (
	"net/http"
)

func start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8000", nil)

}
