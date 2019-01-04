package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//mario := person{"Mario", "Peña", contactInfo{"mario.pc2009@gmail.com", 87452}} //First Way to initialize a struct
	mario := person{ //Second Way to initialize a struct
		firstName: "Mario",
		lastName:  "Peña",
		contact: contactInfo{
			email:   "mario.pc2009@gmail.com",
			zipCode: 87452,
		},
	}
	/*var mario person //Third Way to initialize a struct
	mario.firstName = "Mario"
	mario.lastName = "Peña"
	mario.contact.email = "mario.pc2009@gmail.com"
	mario.contact.zipCode = 87452*/
	mario.print()
	mario.updateName("Julieta")
	mario.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
