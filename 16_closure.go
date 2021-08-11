package main

import "fmt"

func main() {
	name := "aditya"
	counter := 0

	increment := func()  {
		counter++
		fmt.Println("Increment")
		name := "fathur"
		fmt.Println(name)
	}

	increment()
	increment()
	
	fmt.Println(name)
	fmt.Println(counter)
}