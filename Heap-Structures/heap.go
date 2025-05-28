package heap

type heap[T comparable] struct {
	nodes         []T
	orderCriteria func(T, T) bool
}

func heapInit[T comparable](sort func(T, T) bool) *heap[T] {
	heap := &heap[T]{}
	heap.orderCriteria = sort
	return heap
}
func (self *heap[T]) isEmpty() bool {
	if nodes := self.nodes; nodes != nil {
		return len(nodes) == 0
	}
	return true
}

func (self *heap[T]) count() int {
	if !self.isEmpty() {
		return len(self.nodes)
	}
	return 0
}

func (self *heap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

func (self *heap[T]) leftChildIndex(index int) int {
	return 2*index + 1
}

func (self *heap[T]) peek() (T, bool) {
	if self.isEmpty() {
		var element T
		return element, false
	}
	element := self.nodes[0]
	return element, true
}

func (self *heap[T]) insert(value T) {
	self.nodes = append(self.nodes, value)
	self.shiftUp(self.count() - 1)
}

func (self *heap[T]) pop() (T, bool) {
	if self.isEmpty() {
		var value T
		return value, false
	}

	if self.count() == 1 {
		value := self.nodes[0]
		self.nodes = self.nodes[1:]
		return value, true
	} else {
		value := self.nodes[0]
		self.nodes[0] = self.nodes[len(self.nodes)-1]
		self.nodes = self.nodes[:len(self.nodes)-1]
		self.shiftDown(0)
		return value, true
	}
}

func (self *heap[T]) shiftUp(index int) {
	childIndex := index
	child := self.nodes[childIndex]
	parentIndex := self.parentIndex(childIndex)

	for childIndex > 0 && self.orderCriteria(child, self.nodes[parentIndex]) {
		self.nodes[childIndex] = self.nodes[parentIndex]
		childIndex = parentIndex
		parentIndex = self.parentIndex(childIndex)
	}

	self.nodes[childIndex] = child
}

func (self *heap[T]) shiftDown(indicies ...int) {

	if len(indicies) == 1 {
		self.shiftDown(indicies[0], self.count())
		return
	}

	index := indicies[0]
	endIndex := indicies[1]
	leftChildIndex := self.leftChildIndex(index)
	rightChildIndex := leftChildIndex + 1

	first := index
	if leftChildIndex < endIndex && self.orderCriteria(self.nodes[leftChildIndex], self.nodes[first]) {
		first = leftChildIndex
	}
	if rightChildIndex < endIndex && self.orderCriteria(self.nodes[rightChildIndex], self.nodes[first]) {
		first = rightChildIndex
	}

	if first == index {
		return
	}

	self.nodes[index], self.nodes[first] = self.nodes[first], self.nodes[index]
	self.shiftDown(first, endIndex)
}
