package main

import "fmt"

/* Struct
-sebuah template data yang digunakan untuk menggabungkan nol / lebih tipe data lainnya dalam satu kesatuan
-data di struct disimpan dalam field, struct adalah kumpulan dari field
-template data / prototype data
*/

type Person struct {
	Name, Address	string
	Age				int
}

func (person Person) sayHello(people string) {
	fmt.Println("Hello", people + ", My name is", person.Name)
}

func main() {
	var adit Person
	
	adit.Name = "M Aditya FR"
	adit.Address = "Cimahi"
	adit.Age = 23

	fmt.Println(adit)
	fmt.Println(adit.Name)
	fmt.Println(adit.Address)
	fmt.Println(adit.Age)

	fathur := Person{
		Name: "Fathur",
		Address: "Bandung",
	}
	
	fmt.Println(fathur)
	fmt.Println("Name", fathur.Name + ", live in", fathur.Address + ", age", fathur.Age)

	fathur.sayHello("Rahman")
}