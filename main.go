package main

import (
	"fmt"
	dsgo "github.com/fuad7161/dsgo/LinearDataStructures"
)

func main() {
	var dq = dsgo.DequeInt()

	dq.PushFront(100) // 1000
	dq.PushBack(50)   // 100, 50
	dq.PushFront(90)  // 90 , 100 , 50

	fmt.Println(dq.Front())
	fmt.Println(dq.Back())
	dq.PopFront() // 100, 50

	fmt.Println(dq.Front())
	fmt.Println(dq.Back())
}
