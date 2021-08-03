package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type ContactDetailsResponse struct {
	Success bool
	Details ContactDetails
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/contact.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		fmt.Println("Process the give message", details)
		tmpl.Execute(w, ContactDetailsResponse{true, details})
	})
	http.ListenAndServe(":8080", nil)
}
