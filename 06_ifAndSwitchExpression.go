package main

import "fmt"

func main() {
	value := 4
	value2 := 9
	nama := "fathur"

	if value <= value2 {
		fmt.Println(true)
	}

	if nama == "Adit" {
		fmt.Println("Hi, Adit")
	} else if nama == "fathur" {
		fmt.Println("Hi, fathur")
	} else {
		fmt.Println("Siapa kamu?!")
	}

	if length:= 5; length >= 5 {
		fmt.Println("perluas panjang")
	} else {
		fmt.Println("panjang default 5")
	}

	switch nama {
	case "Adit":
		fmt.Println("hi, adit")
	case "fathur":
		fmt.Println("hi, fathur")
	default:
		fmt.Println("siapa kamu?!")
	}

	switch length:= 3; length >= 5 {
	case true:
		fmt.Println("perluas panjang")
	case false:
		fmt.Println("panjang default 5")
	}
	usia := 17
	switch {
	case usia >= 20:
		fmt.Println("Dewasa")
	case usia >= 12:
		fmt.Println("Remaja")
	case usia >= 5:
		fmt.Println("Anak-anak")
	default:
		fmt.Println("Balita")
	}
}