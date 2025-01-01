package main

import "fmt"

type contact_info struct {
	email    string
	postcode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contact_info
}

func main() {
	alex := person{"Alex", "Smith", contact_info{"asmith@domain.uk", "SW11 ER6"}}                                                // Go will infer the order based on the order they are defined in the struct declaration
	bob := person{firstName: "Bob", lastName: "Geldoff", contactInfo: contact_info{email: "bg@domain.uk", postcode: "SW11 ER6"}} // Order does not matter
	var tony person                                                                                                              // tony is intialised with no values for firstName or lastName
	var contactInfo contact_info
	tony.firstName = "Tony"
	tony.lastName = "Montana"
	contactInfo.email = "tm@domain.uk"
	contactInfo.postcode = "SW11 ER6"
	tony.contactInfo = contactInfo

	fmt.Println(alex, bob, tony)
}
