package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		fmt.Fprintf(w, "Hello there!")
	})
	http.ListenAndServe(":8080", nil)
}
