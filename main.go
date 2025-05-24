package main

import (
	"fmt"
	dsgo "github.com/fuad7161/dsgo/LinearDataStructures"
)

type person struct {
	name string
	age  int
}

func main() {
	var vv = dsgo.StackInt()
	vv.Push(100)

	var v = dsgo.NewStack[person]()
	p1 := &person{"fuad", 20}
	p2 := &person{"masud", 22}
	v.Push(*p1)
	v.Push(*p2)
	value1, _ := v.Pop()
	value2, _ := v.Pop()
	fmt.Println(value1)
	fmt.Println(value2)
}
