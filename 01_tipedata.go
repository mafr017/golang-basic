package main

import "fmt"

/* Tipe Data Number
-integer
	-int8, int16, int32, int64, uint8(tidak punya nilai negatif, jadi bergeser ke positif nilai rangenya)
	-byte(unit8), rune(int32), int(int32/64), uint(uint32/64)
-floating point
	-float32, float64, complex64, complex128
*/

/* Tipe Data Boolean
-bool
	-true, false
*/

/* Tipe Data String
-string
	-menggunakan tanda petik dua ""
	-menghitung jumlah char pada string gunakan len("string") => 6 char
*/

/* Variable
-var
	-hanya dapat menyimpan tipe data yang sama
	-dapat diubah-ubah nilainya
	-var dapat diganti dengan :=  mis.(nama := "MAdityaFR")
*/

/* Constant
-const
	-hanya dapat menyimpan tipe data yang sama
	-tidak dapat diubah-ubah nilainya
*/

func main()  {
	// var angka int
	// var nama = "M Aditya FR"
	var (
		angka int
		nama = "M Aditya FR"
	)
	negara := "Indonesia"
	const piNum float32 = 3.14
	const (
		ina = "indonesia"
		pi = 3.14
	)

	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Dua Setengah =", 2.5)

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)
	
	fmt.Println("Nama Saya =", nama)
	// nama = 22 error karna tidak sama tipe datanya
	fmt.Println("Asal =", negara)

	angka = 17
	fmt.Println("angka awal =", angka)
	angka = 21
	fmt.Println("angka akhir =", angka)
	
	fmt.Println("pi number =", piNum)
	// piNum = 21 error
}