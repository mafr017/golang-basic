package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Adit"
	names[1] = "Fathur"
	names[2] = "Rahman"

	var kota = [3]string{
		"cimahi",
		"bandung",
		"jakarta",
	}
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	names[0] = "Aditya"
	fmt.Println(names[0])

	fmt.Println(kota)
	fmt.Println(len(kota))
}