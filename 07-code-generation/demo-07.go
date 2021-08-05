package main

import (
	"fmt"
	"os"
	"text/template"
)

var collectionUtilTemplate = `
type {{.}}s []{{.}}

func (items *{{.}}s) IndexOf(item {{.}}) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *{{.}}s) Includes(item {{.}}) bool {
	return items.IndexOf(item) != -1
}

func (items *{{.}}s) Any(criteria func({{.}}) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items *{{.}}s) All(criteria func({{.}}) bool) bool {
	for _, item := range *items {
		if !criteria(item) {
			return false
		}
	}
	return true
}

func (items *{{.}}s) Filter(criteria func({{.}}) bool) *{{.}}s {
	result := &{{.}}s{}
	for _, item := range *items {
		if criteria(item) {
			*result = append(*result, item)
		}
	}
	return result
}
`

func main() {
	typeName := "Employee"
	file, _ := os.Create(typeName + "s.go")
	defer file.Close()
	parsedT := template.Must(template.New("").Parse(collectionUtilTemplate))
	parsedT.Execute(file, "Employee")
	fmt.Println("Done")
}
