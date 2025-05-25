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
	var s1 = dsgo.StackInt()
	s1.Push(100)
	s1.Push(200)
	var s2 = dsgo.StackInt()
	//var s2 = s1.ToSlice()
	s2.Push(300)
	fmt.Println(s1.Top())
	fmt.Println(s2.Top())
	s1.Swap(s2)
	fmt.Println(s1.Top())
	fmt.Println(s2.Top())
}
