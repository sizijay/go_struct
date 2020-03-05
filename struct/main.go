package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// var bob person
	bob := person{"Bobby", "Anderson", contactInfo{"bob@gmail.com", 7394739}}
	fmt.Println(bob)
	fmt.Printf("%+v \n", bob)
	fmt.Printf("First Name: %v\nPhone: %v\n", bob.firstName, bob.contact.phone)
	fmt.Println()

	bob.print()

	fmt.Println("\nUpdating First Name...")

	bob.updateFName("Ginny")
	bob.print()
}

func (p person) print() {
	fmt.Printf("\n****Printer Working******\n\nFirst Name\t:\t%v\nLastName\t:\t%v\nEmail\t\t:\t%v\nPhone\t\t:\t%v\n\n", p.firstName, p.lastName, p.contact.email, p.contact.phone)
}

func (p *person) updateFName(newFName string) {
	p.firstName = newFName
}
