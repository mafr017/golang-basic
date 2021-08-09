package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter)  {
// func sayHelloWithFilter(name string, filter func(string) string)  {
	namefilter := filter(name)
	fmt.Println("Hello,", namefilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}