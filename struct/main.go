package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	IsActive  bool
	MoreInfo Address
}

type Address struct {
	Country string
}

func main() {
	var Chetan Person

	Chetan.FirstName = "Chetan"

	fmt.Println(Chetan)

	var p2 = Person{
		FirstName: "CHETAN",
		LastName: "RAJAWAT",
		Age: 10,
	}

	fmt.Println(p2);


	p3 := new(Person);

	p3.IsActive = false;
	p3.MoreInfo.Country = "JAPAN"


	fmt.Println(p3);
}
