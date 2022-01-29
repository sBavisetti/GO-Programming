package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	gender    string
	age       int
	contact   contactInfo
}

func main() {

	// declaring and initializing a struct
	s := person{firstName: "Reed", lastName: "Richards", gender: "Male", age: 44}

	// displaying the all the values in a struct
	s.print()

	// updating or editing the values in a struct and struct within struct
	s.firstName = "Genna"
	s.contact.email = "genna@marvel.com"
	s.contact.zip = 12345

	fmt.Printf("%+v\n", s)

	// declaring a variable and printing to see zero values
	var p person
	fmt.Println("p", p)

	// declaring a struct within a struct
	v := person{"Barry", "Allen", "Male", 26, contactInfo{"BarryAllen@starlabs.com", 223345}}
	fmt.Printf("%+v\n", v)

	//second way of declaring a variable
	arthur := person{
		firstName: "Arthur",
		lastName:  "PennDragon",
		gender:    "Male",
		age:       25,
		contact: contactInfo{
			email: "arthur.penndragon@camlot.com",
			zip:   90525,
		},
	}

	fmt.Printf("%+v\n", arthur)

	arthur.updateName("Urther")
	arthur.print()

	//declaring a variable for woring on pointers
	matt := person{
		firstName: "Matt",
		lastName:  "Murdock",
		gender:    "Male",
		age:       25,
		contact: contactInfo{
			email: "daredevil@hellskitchen.com",
			zip:   11015,
		},
	}

	// '&' operator -> gives the RAM memory address to where the variable is pointing
	mattPointer := &matt
	mattPointer.updateLastName("Cooper")
	matt.print()

}

func (p person) print() {
	fmt.Println("First name : ", p.firstName)
	fmt.Println("Last name  : ", p.lastName)
	fmt.Println("Gender     : ", p.gender)
	fmt.Println("Age        : ", p.age)
	fmt.Println("Email      : ", p.contact.email)
	fmt.Println("Zip        : ", p.contact.zip)

}

func (p *person) updateName(newfirstName string) {
	p.firstName = newfirstName
}

// '*person' -> This is a type description - It means we are working with a pointer to a person
func (pointerToPerson *person) updateLastName(newLastName string) {
	// '*' operator -> gives the values present in the memory address
	(*pointerToPerson).lastName = newLastName
}
