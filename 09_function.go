package main

import (
	"fmt"
	"strconv"
)

func main() {
	sayHello("Adit")
	sayHello("Go-Lang")
	
	sayHelloFull("Aditya", "Fathur")
	
	fmt.Println("Bilangan ini", checkGanjilGenap(25))

	ket1, ket2 := identitas("Aditya", 23)
	ket3, _ := identitas("Fathur", 18)
	fmt.Println(ket1)
	fmt.Println(ket2)
	fmt.Println(ket3)
}

func sayHello(name string)  {
	fmt.Println("Hello,", name)
}

func sayHelloFull(name string, lastName string)  {
	fmt.Println("Hello,", name, lastName)
}

func checkGanjilGenap(val int) string {
	var hasil string
	if val % 2 == 0 {
		hasil = strconv.Itoa(val) + " adalah genap"
	} else {
		hasil = strconv.Itoa(val) + " adalah ganjil"
	}
	return hasil
}

func identitas(name string, umur int) (string, string)  {
	var keterangan1 string
	var keterangan2 string
	
	keterangan1 = "Hallo, " + name
	
	switch {
	case umur >= 20:
		keterangan2 = "Dewasa"
	case umur >= 12:
		keterangan2 = "Remaja"
	case umur >= 5:
		keterangan2 = "Anak-anak"
	default:
		keterangan2 = "Balita"
	}
	
	return keterangan1, keterangan2
}