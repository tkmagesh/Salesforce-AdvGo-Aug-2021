package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var templateStr = "Hi {{.}} \n"
	te := template.New("greeter")
	t, err := te.Parse(templateStr)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(os.Stdout, "Gopher")
}
