package main

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	Data string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Data)
}

func Hello() Handler {
	return Handler{"Hello"}
}

func String(s string) Handler {
	return Handler{s}
}

// http://golang.org/pkg/net/http/
func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/", Hello())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
