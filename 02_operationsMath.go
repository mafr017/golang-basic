package main

import (
	"fmt"
)

func main()  {
	var a = 7
	var b = 4
	var c = 21
	var result = a + b
	fmt.Println(result)
	
	result = a - b
	fmt.Println(result)
	fmt.Println(a - c)

	c /= a
	fmt.Println(c)

	c++
	fmt.Println(c)

	var negatif = -c
	fmt.Println(negatif)

	var nilaitrue = true
	var negasi = !nilaitrue
	fmt.Println(negasi)

	var nilaifalse = false
	var hasil bool = nilaitrue == nilaifalse
	fmt.Println(hasil)

	var nama1 = "adit"
	var nama2 = "Adit"
	var nama3 = "Adit"
	hasil = nama1 == nama2
	fmt.Println(hasil)
	fmt.Println(nama1 != nama3)
	
	fmt.Println(100 > 65)
	fmt.Println(777 <= 100)
	
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}