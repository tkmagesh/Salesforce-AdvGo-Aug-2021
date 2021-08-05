package main

import (
	"log"
	"os"
	"text/template"
)

type Name struct {
	First string
	Last  string
}

func main() {
	var templateStr = "Hi {{.First}} {{.Last}} \n"
	te := template.New("greeter")
	t, err := te.Parse(templateStr)
	if err != nil {
		log.Fatalln(err)
	}
	user := Name{First: "Magesh", Last: "Kuppan"}
	err = t.Execute(os.Stdout, user)
}
