package main

import "fmt"

/* Nil
-data kosong
-hanya bisa digunakan di beberapa tipe data interface, function, map, slice, pointer, dan channel
*/

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}
}

func main() {
	fmt.Println(newMap(""))
	fmt.Println(newMap("adit"))
}