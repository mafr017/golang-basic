package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked,", name)
	} else {
		fmt.Println("You're Welcome,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anonymous"
	}

	registerUser("Anonymous", blacklist)
	registerUser("Adit", blacklist)
	
}