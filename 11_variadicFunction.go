package main

import "fmt"

func main() {
	fmt.Println(countAll("+", 1,2,3,4,5,6,7,8,9))
	numbers := []int{100,1,2,3,4,5,6,7,8,9}
	fmt.Println(countAll("-", numbers...))
}

func countAll(tipe string, angka ...int) int {

	total := 0
	
	for i, v := range angka {
		if i == 0 {
			total = v
			continue
		}
		if tipe == "+" {
			total += v
		} else if tipe == "-" {
			total -= v
		} else {
			total = 666
		}
	}
	
	return total
}