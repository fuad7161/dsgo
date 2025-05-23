package main

import (
	"fmt"
	dsgo "main/LinearDataStructures"
)

func main() {
	var q = dsgo.NewQueue[int]()
	q.Push(100)
	val, _ := q.Pop()
	fmt.Println(val)
}
