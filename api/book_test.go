package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.L Reimer", ISBN: "0987654321"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.L Reimer","isbn":"0987654321"}`, string(json), "Book JSON marshalling worng")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.L Reimer","isbn":"0987654321"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.L Reimer", ISBN: "0987654321"}, book, "implement me.")

}
