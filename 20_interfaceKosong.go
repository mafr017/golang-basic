package main

import "fmt"

// Interface kosong mirip dengan any di kotlin

func apaSaja(pilih int) interface{} {
	if pilih == 1 {
		return "string"
	} else if pilih == 2 {
		return 1234567890
	} else if pilih == 3 {
		return 3.14
	} else {
		return []string{"qwert", "yuiop", "asdfg", "hjkl"}
	}
}

func main() {
	// var data int = apaSaja(1) 	//tidak bisa error karna type data
	var data interface{} = apaSaja(9)
	data2 := apaSaja(1)

	fmt.Println(data)
	fmt.Println(data2)
	
	data2 = apaSaja(3)
	fmt.Println(data2)
	if data2 == 3.14 {
		fmt.Println(data2, "adalah angka pi")
	}
}