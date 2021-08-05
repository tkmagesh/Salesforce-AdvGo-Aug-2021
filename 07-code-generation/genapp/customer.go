package main

//go:generate go run gen-col.go -N "Customer" -P "domain"
type Customer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
