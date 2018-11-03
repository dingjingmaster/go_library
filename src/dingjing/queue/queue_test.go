package queue

import "fmt"

func test1() {
	list := New()

	fmt.Println(list.Size())
	list.Push("sss")
	list.Push("bbb")
	list.Push("ccc")
	fmt.Println(list.Size())
	fmt.Println(list.Pop())
	fmt.Println(list.Size())
	fmt.Println(list.Pop())
	fmt.Println(list.Size())
	fmt.Println(list.Pop())
	fmt.Println(list.Size())
}
