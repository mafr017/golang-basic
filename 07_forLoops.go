package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Index -", i)
	}

	warna := [...]string{"merah", "hijau", "kuning", "biru", "oranye", "putih", "hitam"}
	fmt.Println(len(warna))

	for _, element := range warna {
		// fmt.Println("index", index, "=", element)
		fmt.Println(element)
	}

	person := map[string]string{
		"name": "adit",
		"address": "cimahi",
	}

	for i, v := range person {
		fmt.Println(i, "=", v)
	}
}