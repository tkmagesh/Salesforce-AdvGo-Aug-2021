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

type Users struct {
	Names []Name
}

func main() {
	/* 	var templateStr = `
	{{with .Names}}
	{{range .}}
	Hi {{.First}} {{.Last}}
	{{end}}
	{{end}}
		` */

	var templateStr = `
{{range .Names}}	
	{{$msg}} {{.First}} {{.Last}}
{{end}}
	`
	names := []Name{
		{First: "Joe", Last: "Smith", TimeOfDay: "morning"},
		{First: "Jane", Last: "Doe", TimeOfDay: "afternoon"},
		{First: "John", Last: "Doe", TimeOfDay: "evening"},
	}
	users := Users{Names: names}
	parsedT := template.Must(template.New("").Parse(templateStr))
	parsedT.Execute(os.Stdout, users)
}
