package dsgo

import "fmt"

// Heap This is max heap implementation
type Heap struct {
	arr []int
}

func (h *Heap) Insert(val int) {
	h.arr = append(h.arr, val)
	if len(h.arr) == 1 {
		return
	}
	updateAfterInsert(h.arr, len(h.arr)-1)
}

func (h *Heap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}
	var element = h.arr[0]
	h.arr = h.arr[1:]
	updateAfterRemove(h)
	return element
}
func updateAfterRemove(h *Heap) {
	for i := 0; i < len(h.arr); i++ {
		if leftChild(i) < len(h.arr) && h.arr[i] <= h.arr[leftChild(i)] {
			h.arr[i], h.arr[leftChild(i)] = h.arr[leftChild(i)], h.arr[i]
			i = leftChild(i)
		} else if rightChild(i) < len(h.arr) && h.arr[i] < h.arr[rightChild(i)] {
			h.arr[i], h.arr[rightChild(i)] = h.arr[rightChild(i)], h.arr[i]
			i = rightChild(i)
		} else {
			return
		}
	}
}

func (h *Heap) ShowArray() {
	for i := 0; i < len(h.arr); i++ {
		fmt.Printf("%v ", h.arr[i])
	}
	fmt.Println()
}

func updateAfterInsert(arr []int, idx int) {
	for idx > 0 {
		if arr[idx] > arr[parent(idx)] {
			arr[idx], arr[parent(idx)] = arr[parent(idx)], arr[idx]
		}
		idx = parent(idx)
	}
}

func parent(index int) int {
	return ((index + 1) / 2) - 1
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}
