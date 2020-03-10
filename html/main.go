package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	FName string
	LName string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"f3": first3,
}

func init() {
	// tpl = template.Must(template.ParseFiles("1.gohtml"))

	//for passing fnctions to the template
	tpl = template.Must(template.New("1.gohtml").Funcs(fm).ParseFiles("1.gohtml"))
}

func first3(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}

func main() {
	//people := []string{"one", "two", "three", "four", "five"}

	// err := tpl.Execute(os.Stdout, counts)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//using the structs
	p1 := person{
		FName: "Joseph",
		LName: "Dilan",
	}

	p2 := person{
		FName: "Lahiru",
		LName: "Nadeemal",
	}

	p3 := person{
		FName: "Kathy",
		LName: "Perera",
	}

	p4 := person{
		FName: "Ginny",
		LName: "Weasely",
	}

	p5 := person{
		FName: "Sampath",
		LName: "Fernando",
	}

	people := []person{p1, p2, p3, p4, p5}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatalln(err)
	}

	//two data types can also be passed and can be accessed with the range as well

}
