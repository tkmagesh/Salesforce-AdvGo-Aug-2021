package main

import (
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
	parsedT := template.Must(template.New("name").Parse(templateStr))
	//user := Name{First: "Magesh", Last: "Kuppan", TimeOfDay: "Morning"}
	user := Name{First: "Magesh", Last: "Kuppan"}
	parsedT.Execute(os.Stdout, user)
}
