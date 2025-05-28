package main

import (
	"fmt"
	heap "github.com/fuad7161/dsgo/Heap-Structures"
)

func main() {
	var pq = heap.PriorityQueue(heap.MaxString())

	pq.Push("Fuad")
	pq.Push("Masud")

	fmt.Println(pq.Top())
}
