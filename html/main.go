package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	counts := []string{"one", "two", "three", "four", "five"}

	err := tpl.Execute(os.Stdout, counts)

	if err != nil {
		log.Fatalln(err)
	}
}
