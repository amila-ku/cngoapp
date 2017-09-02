package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Name, Autho and ISBN
// adding `json:title specifies go to marshall Title as lowercase in jason`
type Book struct {
	// defines the book
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
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

//Books slice of all books
var Books = []Book{Book{Title: "Loosing Virginity", Author: "Richard Branson", ISBN: "0777654321"}, Book{Title: "Screw It, Lets do IT", Author: "Rischard Branson", ISBN: "9987654321"}}

// BooksHandleFunc to be used as http.HandleFunch for Book api

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.Write(b)

}
