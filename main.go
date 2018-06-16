// Define the current package for import/export purposes.
package main

import (
	_ "encoding/json"
	"github.com/gorilla/mux"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

// Book Struct (Model).
type Book struct {
	// <field name> <field type> <JSON field mapping>
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"` // *Author points to another struct.
}

// Author Struct.
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Get All Books.
//	All route handlers must accept ResponseWriter and Request parameters.
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get Single Book.
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create Book.
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update Book.
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book.
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

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
