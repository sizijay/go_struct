package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// var bob person
	bob := person{"Bobby", "Anderson"}
	fmt.Println(bob)
	fmt.Printf("%+v \n", bob)
	fmt.Printf("First Name: %v", bob.firstName)
	fmt.Println()
}
