package heap

type priorityQueue[T comparable] struct {
	heap *heap[T]
}

// Max Functions
func MaxInt() func(int, int) bool {
	return func(a, b int) bool { return a > b }
}

func MaxInt8() func(int8, int8) bool {
	return func(a, b int8) bool { return a > b }
}

func MaxInt16() func(int16, int16) bool {
	return func(a, b int16) bool { return a > b }
}

func MaxInt32() func(int32, int32) bool {
	return func(a, b int32) bool { return a > b }
}

func MaxInt64() func(int64, int64) bool {
	return func(a, b int64) bool { return a > b }
}

func MaxUint() func(uint, uint) bool {
	return func(a, b uint) bool { return a > b }
}

func MaxUint8() func(uint8, uint8) bool {
	return func(a, b uint8) bool { return a > b }
}

func MaxUint16() func(uint16, uint16) bool {
	return func(a, b uint16) bool { return a > b }
}

func MaxUint32() func(uint32, uint32) bool {
	return func(a, b uint32) bool { return a > b }
}

func MaxUint64() func(uint64, uint64) bool {
	return func(a, b uint64) bool { return a > b }
}

func MaxFloat32() func(float32, float32) bool {
	return func(a, b float32) bool { return a > b }
}

func MaxFloat64() func(float64, float64) bool {
	return func(a, b float64) bool { return a > b }
}

func MaxString() func(string, string) bool {
	return func(a, b string) bool { return a > b }
}

// Min Functions
func MinInt() func(int, int) bool {
	return func(a, b int) bool { return a < b }
}

func MinInt8() func(int8, int8) bool {
	return func(a, b int8) bool { return a < b }
}

func MinInt16() func(int16, int16) bool {
	return func(a, b int16) bool { return a < b }
}

func MinInt32() func(int32, int32) bool {
	return func(a, b int32) bool { return a < b }
}

func MinInt64() func(int64, int64) bool {
	return func(a, b int64) bool { return a < b }
}

func MinUint() func(uint, uint) bool {
	return func(a, b uint) bool { return a < b }
}

func MinUint8() func(uint8, uint8) bool {
	return func(a, b uint8) bool { return a < b }
}

func MinUint16() func(uint16, uint16) bool {
	return func(a, b uint16) bool { return a < b }
}

func MinUint32() func(uint32, uint32) bool {
	return func(a, b uint32) bool { return a < b }
}

func MinUint64() func(uint64, uint64) bool {
	return func(a, b uint64) bool { return a < b }
}

func MinFloat32() func(float32, float32) bool {
	return func(a, b float32) bool { return a < b }
}

func MinFloat64() func(float64, float64) bool {
	return func(a, b float64) bool { return a < b }
}

func MinString() func(string, string) bool {
	return func(a, b string) bool { return a < b }
}

func PriorityQueue[T comparable](sort func(T, T) bool) *priorityQueue[T] {
	pQueue := &priorityQueue[T]{heap: heapInit(sort)}
	return pQueue
}

func (self *priorityQueue[T]) Empty() bool {
	return self.heap.isEmpty()
}

func (self *priorityQueue[T]) Size() int {
	return self.heap.count()
}

func (self *priorityQueue[T]) Top() (T, bool) {
	return self.heap.peek()
}

func (self *priorityQueue[T]) Push(element T) {
	self.heap.insert(element)
}

func (self *priorityQueue[T]) Pop() (T, bool) {
	return self.heap.pop()
}

func (self *priorityQueue[T]) Swap(pq *priorityQueue[T]) {
	*self, *pq = *pq, *self
}
