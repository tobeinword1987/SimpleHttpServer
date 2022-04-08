package main

import (
	"net/http"
)

func main() {
	// Create a basic http
	http.Handle("/users", http.HandlerFunc(ExampleHandler))
	http.ListenAndServe(":3333", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(getUsersFromFile())
	}
}
