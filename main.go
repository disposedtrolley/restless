// Define the current package for import/export purposes.
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	// Initialise the Mux router.
	// 	:= creates new variables using type inference.
	r := mux.NewRouter()

	// Route Handlers / Endpoints.
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run the server.
	// Wrapped in log.Fatal() so we get an error if it fails.
	log.Fatal(http.ListenAndServe(":8000", r))
}
