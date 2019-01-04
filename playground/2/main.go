package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mySlice := []string{"Hola", "mundo"}

	person := person{firstName: "Mario", lastName: "Peña"}

	updateSlice(mySlice)
	updatePerson(person)

	fmt.Println(mySlice)
	fmt.Println(person)
}

func updateSlice(s []string) {
	s[0] = "Chao"
}

func updatePerson(p person) {
	p.firstName = "Julieta"
}
