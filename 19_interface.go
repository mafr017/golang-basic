package main

import "fmt"

/* Interface
-tipe data Abstract, tidak memiliki implementasi langsung
-berisikan definisi-definisi method
-digunakan untuk kontrak
*/

type HasName interface {
	GetName() string 
}

func SayHai(hasName HasName)  {
	fmt.Println("Hai", hasName.GetName())
}

type Person struct {
	Name, Address	string
	Age				int
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var adit Person
	adit.Name = "Aditya"

	SayHai(adit)

	var kucing Animal
	kucing.Name = "Meng"

	SayHai(kucing)
}