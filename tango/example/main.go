package main

import (
	"fmt"
	"github.com/ikbear/tango"
	"net/http"
	"net/url"
)

type Book struct {
}

func (book Book) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}
func (book Book) Post(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}
func (book Book) Put(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}

func main() {
	port := 3000
	book := new(Book)
	app := tango.NewAPI()
	app.Register(book, "/books")
	fmt.Println("Listen on port: ", port)
	app.Start(port)
}
