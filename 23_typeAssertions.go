package main

import "fmt"

/* Type Assertions
-kemampuan merubah tipe data menjadi tipe data yang diinginkan
-digunakan saat bertemu dengan interface kosong
*/

func apaSaja(pilih int) interface{} {
	if pilih == 1 {
		return "string"
	} else if pilih == 2 {
		return 1234567890
	} else if pilih == 3 {
		return 3.14
	} else if pilih == 4 {
		return false
	} else {
		return []string{"qwert", "yuiop", "asdfg", "hjkl"}
	}
}

func main() {
	var iniString string
	var iniInt int

	result := apaSaja(1)
	resultString := result.(string)
	iniString = resultString
	fmt.Println(iniString)

	result = apaSaja(2)
	resultInt := result.(int)
	iniInt = resultInt
	fmt.Println(iniInt)

	result = apaSaja(4)
	// resultAnother := result.(int)	//ini akan error
	// fmt.Println("ini akan error",  resultAnother)
	switch value := result.(type) {
	case string:
		fmt.Println("ini string", value)
	case int:
		fmt.Println("ini int", value)
	case float32:
		fmt.Println("ini float", value)
	case float64:
		fmt.Println("ini float", value)
	case bool:
		fmt.Println("ini boolean", value)
	default:
		fmt.Println("something else")
	}


}