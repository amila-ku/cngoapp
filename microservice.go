package main

import (
	"fmt"
	"net/http"
	"os"

	"./api"
)

func main() {
	fmt.Println("Starting HTTP listener on " + port())

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go")
}

func echo(w http.ResponseWriter, r *http.Request) {
	messege := r.URL.Query()["messege"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, messege)
}
