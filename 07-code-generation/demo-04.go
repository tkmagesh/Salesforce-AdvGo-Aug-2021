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
{{range .}}
Hi {{.First}} {{.Last}}
{{end}}
	`
	names := []Name{
		{First: "Joe", Last: "Smith", TimeOfDay: "morning"},
		{First: "Jane", Last: "Doe", TimeOfDay: "afternoon"},
		{First: "John", Last: "Doe", TimeOfDay: "evening"},
	}
	parsedT := template.Must(template.New("").Parse(templateStr))
	parsedT.Execute(os.Stdout, names)
}
