package dsgo

// Heap This is max heap implementation
type heap struct {
	arr []int
}

// NewHeap creating new Heap
func NewHeap() heap {
	return heap{}
}

// Insert : Insert element into heap
func (h *heap) Insert(val int) {
	h.arr = append(h.arr, val)
	if len(h.arr) == 1 {
		return
	}
	updateAfterInsert(h.arr, len(h.arr)-1)
}

// Pop : Pop minimum element from the heap
func (h *heap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}
	var element = h.arr[0]
	h.arr = h.arr[1:]
	updateAfterRemove(h)
	return element
}

// updateAfterRemove after removing root element update the heap tree
func updateAfterRemove(h *heap) {
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

// updateAfterInsert after insert new element update the min heap
func updateAfterInsert(arr []int, idx int) {
	for idx > 0 {
		if arr[idx] > arr[parent(idx)] {
			arr[idx], arr[parent(idx)] = arr[parent(idx)], arr[idx]
		}
		idx = parent(idx)
	}
}

// parent return the parent index from heap array
func parent(index int) int {
	return ((index + 1) / 2) - 1
}

// leftChild return the left index of an element from heap array
func leftChild(index int) int {
	return index*2 + 1
}

// rightChild return the right index of an element from heap array
func rightChild(index int) int {
	return index*2 + 2
}
