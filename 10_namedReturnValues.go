package main

import "fmt"

func main() {
	firstName, middleName, lastName, _ := getName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}

func getName() (firstName, middleName, lastName string, age int) {
	firstName = "M Aditya"
	middleName = "Fathur"
	lastName = "Rahman"
	age = 23
	return
}