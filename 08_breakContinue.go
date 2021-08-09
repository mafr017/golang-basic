package main

import "fmt"

func main() {
	fmt.Println("Bilangan genap")
	for i := 1; i < 100; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		if i >= 6 {
			break
		}
		fmt.Println(i)
	}
}