package main

import "fmt"

/* POINTER
- Pass by Value
	-secara default di Go-Lang, semua variable itu di passing by value, bukan by reference
	-ketika mengirim sebuah variable ke function/method/variable lain sebenarnya yang dikirim adalah duplikasi valuenya
-Pointer
	-kemampuan membuat reference ke lokasi data di memory yang sama, tanpa, menduplikasi data yang sudah ada
	-bisa membuat pass by reference
	-untuk menggunakan pointer bisa gunakan operator '&' diikuti nama variabelnya
-Operator *
	-saat mengubah variable pointer, maka yang berubah hanya variable tersebut
	-jika ingin mengubah seluruh variable yang mengacu ke data tersebut, bisa gunakan operator '*'
-new()
	-function untuk membuat data pointer baru dengan data yang kosong (tidak ada data awal)
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Cimahi", "JaBar", "Indonesia"}
	address2 := address1
	address3 := &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address1)

	address3.City = "Padalarang"
	fmt.Println(address3)
	fmt.Println(address1)
	
	// address3 = &Address{"Malang", "JaTim", "Indonesia"}
	// fmt.Println(address3)
	// fmt.Println(address1)

	*address3 = Address{"Malang", "JaTim", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address1)
	fmt.Println(address2)

	address4 := new(Address)
	address5 := address4
	fmt.Println(address4)
	fmt.Println(address5)
	address5.City = "New York"
	fmt.Println(address4)
	fmt.Println(address5)
}