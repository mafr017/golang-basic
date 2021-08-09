package main

import "fmt"

func main() {
	sayHALO := sayHello
	fmt.Println(sayHALO("adit"))
}

func sayHello(nama string) string {
	return "Hello, " + nama
}