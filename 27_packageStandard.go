package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	fmt.Println("Ini slice of arg:", args)

	//konversi tipe data
	angkaFloat := 3.14
	tulisan := strconv.FormatFloat(angkaFloat, 'f', -1, 32)
	fmt.Println("ini angka dalam betuk string " + tulisan)

	merdeka := "170845"
	konversi, err := strconv.Atoi(merdeka)
	if err == nil {
		fmt.Println(konversi)
	}

	koversiString := strconv.Itoa(konversi)
	fmt.Println(koversiString)

	//math package
	ni1 := 6
	ni2 := 3.4
	fmt.Println(math.Round(ni2))
	fmt.Println(math.Floor(ni2))
	fmt.Println(math.Ceil(ni2))
	fmt.Println(math.Max(float64(ni1), ni2))

	//time package
	now := time.Now()
	fmt.Println("Sekarang", now)
	now2 := time.Now().Day()
	fmt.Println("Sekarang", now2)

	layout := "2006-01-02"
	parses, _ := time.Parse(layout, "2021-08-14")
	fmt.Println(parses)
	parsee, _ := time.Parse(time.RFC3339, "2021-08-16T13:51:05Z")
	fmt.Println(parsee)


	//regexp
	regex := regexp.MustCompile(`@`)
	bool1 := regex.MatchString("madityafr3@gmail.com")
	fmt.Println(bool1)
	bool2 := regex.MatchString("madityafr3gmail.com")
	fmt.Println(bool2)
}