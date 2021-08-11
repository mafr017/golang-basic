package main

import "fmt"

func HitungFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * HitungFactorial(value-1)
	}
}

func main() {
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(HitungFactorial(5))
}