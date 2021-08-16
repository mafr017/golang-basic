package main

import (
	"container/list"
	"fmt"
)

/* Package Container List
- implementasi double linked list
- data berukuran dynamic
*/

func main() {
	data := list.New()
	data.PushBack("Adit")
	data.PushBack("Fathur")
	data.PushFront("Muhammad")
	data.PushBack("Rahman")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Prev().Value)
	fmt.Println(data.Back().Value)

	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}