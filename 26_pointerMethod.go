package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married() {
func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	fathur := Man{"Fathur"}
	fmt.Println(fathur.Name)
	fathur.Married()
	fmt.Println(fathur.Name)
}