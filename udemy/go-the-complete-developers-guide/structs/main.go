package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Smith"}                      // Go will infer the order based on the order they are defined in the struct declaration
	bob := person{firstName: "Bob", lastName: "Geldoff"} // Order does not matter
	var tony person                                      // tony is intialised with no values for firstName or lastName
	tony.firstName = "Tony"
	tony.lastName = "Montana"

	fmt.Println(alex, bob, tony)
}
