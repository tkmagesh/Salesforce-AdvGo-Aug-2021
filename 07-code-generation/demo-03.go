package main

import (
	"log"
	"os"
	"text/template"
)

type Name struct {
	First     string
	Last      string
	TimeOfDay string
}

func main() {
	var templateStr = `
Hi {{.First}} {{.Last}}, 
{{if .TimeOfDay}} Good Morning! 
{{else}} Have a good day!
{{end}}
	`
	te := template.New("greeter")
	t, err := te.Parse(templateStr)
	if err != nil {
		log.Fatalln(err)
	}
	//user := Name{First: "Magesh", Last: "Kuppan", TimeOfDay: "Morning"}
	user := Name{First: "Magesh", Last: "Kuppan"}
	err = t.Execute(os.Stdout, user)
}
