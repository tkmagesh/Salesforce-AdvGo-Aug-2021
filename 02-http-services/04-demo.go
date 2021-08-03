package main

import (
	"fmt"
	"net/http"
)

func logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logger(foo))
	http.HandleFunc("/bar", logger(bar))
	http.ListenAndServe(":8080", nil)
}
