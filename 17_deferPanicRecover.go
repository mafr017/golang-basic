package main

import "fmt"

/* Defer
-function yang bisa kita jadwalkan untuk dieksekusi meskipun terjadi error di function yang dieksekusi
*/
/* Panic
-function yang akan menghentikan program ketika terjadi error pada saat program kita berjalan (defer tetap akan dieksekusi)
*/
/* Recover
-function yang bisa kita gunakan untuk menangkap data panic, proses panic akan terhenti, sehingga program akan tetap berjalan.
*/
//defer code example
func logging()  {
	fmt.Println("Selesai memanggil function")
	fmt.Println("")
}

func runApplication(value int) {
	defer logging()
	result := 23 / value
	fmt.Println("Run Application")
	fmt.Println(result)
}


//panic & recover code example
func endApp()  {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Ada error di:", message)
	}
}

func runApp(err bool)  {
	defer endApp()
	if err {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApplication(0)

	runApp(true)
	// runApp(false)
	fmt.Println("Adit")
}