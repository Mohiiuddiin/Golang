package main

import "fmt"

type Contact struct {
	email   string
	zipCode string
}

type Person struct {
	firstName   string
	lastName    string
	//contactInfo contact
	Contact
}

func main() {
	person1 := Person{
		firstName: "Mohammed",
		lastName:  "Mohiuddin",
		// contactInfo: contact{
		// 	email:   "mohiuddin@gmail.com",
		// 	zipCode: "4208",
		// },
		Contact : Contact {
			email:   "mohiuddin@gmail.com",
		 	zipCode: "4208",
		},
	}

	(&person1).UpdateName("Mohiuddin","Mohammed")
	//fmt.Println(person1)
	person1.print()
}

func (pointerToPerson *Person) UpdateName(fName,lName string){
	pointerToPerson.firstName = fName
	pointerToPerson.lastName = lName	
}

func (p Person) print(){
	fmt.Printf("%+v",p)
}