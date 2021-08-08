package main

import "fmt"

/* Tipe Data Slice
-potongan dari data array
-ukuran slice bisa berubah
-tipe data slice ada 3
	- pointer = penunjuk data pertama di array para slice
	- length = panjang dari slice
	- capacity = kapasitas slice, dimana length tidak boleh lebih dari capacity
*/

func main() {

	warna := [...]string{"merah", "hijau", "kuning", "biru", "oranye", "putih", "hitam"}
	slice := warna[3:5]
	slice2 := warna[4:]

	fmt.Println(warna)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(warna[:4])
	fmt.Println(warna[5:])

	fmt.Println(append(slice2, "Ungu"))

	slice[1] = "Oren"
	fmt.Println(slice)

	//make([]type, length, capacity)
	newSlice := make([]string, 5, 7)

	newSlice[0] = "lembar1"
	newSlice[1] = "lembar2"
	newSlice[4] = "lembar3"
	fmt.Println(newSlice)

	copySlice := make([]string, 2, 5)
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
