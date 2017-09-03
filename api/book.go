package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Book type with Name, Autho and ISBN
// adding `json:title specifies go to marshall Title as lowercase in jason`
type Book struct {
	// defines the book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

//Books slice of all books
var books = map[string]Book{
	"0777654321": Book{Title: "Loosing Virginity", Author: "Richard Branson", ISBN: "0777654321"},
	"9987654321": Book{Title: "Screw It, Lets do IT", Author: "Rischard Branson", ISBN: "9987654321"},
}

// ToJSON to be used for marshalling of book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

//Write book as JSON
func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.Write(b)

}

// BooksHandleFunc to be used as http.HandleFunch for Book api

/*func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.Write(b)

}
*/

func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// BookHandleFunc to be used as http.HandleFunch for Book api
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// BooksHandleFunc to be used as http.HandleFunch for Book api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	//return all books

	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Unsupported request method."))
	}

}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[book.ISBN] = book
	}
	return exists
}

func DeleteBook(isbn string) {
	_, exists := books[isbn]
	if exists {
		delete(books, isbn)
	}
}
