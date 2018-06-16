// Define the current package for import/export purposes.
package main

import (
	"encoding/json"
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

// Initialise books var as a slice Book struct.
// A slice is a variable-length array.
var books []Book

// Get All Books.
//	All route handlers must accept ResponseWriter and Request parameters.
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode the books array as JSON and write the response.
	json.NewEncoder(w).Encode(books)
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

	// Mock Data - @todo - implement DB.
	books = append(books, Book{ID: "1", ISBN: "9780140186390", Title: "East of Eden", Author: &Author{FirstName: "John", LastName: "Steinbeck"}})
	books = append(books, Book{ID: "2", ISBN: "9782915629217", Title: "Fahrenheit 451", Author: &Author{FirstName: "Ray", LastName: "Bradbury"}})
	books = append(books, Book{ID: "3", ISBN: "9781471331435", Title: "1984", Author: &Author{FirstName: "George", LastName: "Orwell"}})

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
