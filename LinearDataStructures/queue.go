package dsgo

import "fmt"

// Queue is a generic queue implementation using a slice.
type queue[T any] struct {
	elements []T
}

func QueueInt() *queue[int] {
	return &queue[int]{elements: make([]int, 0)}
}

func QueueInt8() *queue[int8] {
	return &queue[int8]{elements: make([]int8, 0)}
}

func QueueInt16() *queue[int16] {
	return &queue[int16]{elements: make([]int16, 0)}
}

func QueueInt32() *queue[int32] {
	return &queue[int32]{elements: make([]int32, 0)}
}

func QueueInt64() *queue[int64] {
	return &queue[int64]{elements: make([]int64, 0)}
}

func QueueUint() *queue[uint] {
	return &queue[uint]{elements: make([]uint, 0)}
}

func QueueUint8() *queue[uint8] {
	return &queue[uint8]{elements: make([]uint8, 0)}
}

func QueueUint16() *queue[uint16] {
	return &queue[uint16]{elements: make([]uint16, 0)}
}

func QueueUint32() *queue[uint32] {
	return &queue[uint32]{elements: make([]uint32, 0)}
}

func QueueUint64() *queue[uint64] {
	return &queue[uint64]{elements: make([]uint64, 0)}
}

func QueueFloat32() *queue[float32] {
	return &queue[float32]{elements: make([]float32, 0)}
}

func QueueFloat64() *queue[float64] {
	return &queue[float64]{elements: make([]float64, 0)}
}

func QueueComplex64() *queue[complex64] {
	return &queue[complex64]{elements: make([]complex64, 0)}
}

func QueueComplex128() *queue[complex128] {
	return &queue[complex128]{elements: make([]complex128, 0)}
}

func QueueString() *queue[string] {
	return &queue[string]{elements: make([]string, 0)}
}

func QueueBool() *queue[bool] {
	return &queue[bool]{elements: make([]bool, 0)}
}

func QueueByte() *queue[byte] {
	return &queue[byte]{elements: make([]byte, 0)}
}

func QueueRune() *queue[rune] {
	return &queue[rune]{elements: make([]rune, 0)}
}

// NewQueue creates and returns a new empty Queue.
func NewQueue[T any]() *queue[T] {
	return &queue[T]{make([]T, 0)}
}

// Empty checks if the queue has no elements.
func (q *queue[T]) Empty() bool {
	return q.Size() == 0
}

// Size returns the number of elements in the queue.
func (q *queue[T]) Size() int {
	return len(q.elements)
}

// Front returns the front element without removing it.
// Returns an error if the queue is empty.
func (q *queue[T]) Front() (T, error) {
	if q.Empty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

// Back returns the front element without removing it.
// Returns an error if the queue is empty.
func (q *queue[T]) Back() (T, error) {
	if q.Empty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.elements[q.Size()-1], nil
}

// Push adds an element to the back of the queue.
func (q *queue[T]) Push(v T) {
	q.elements = append(q.elements, v)
}

// Pop removes and returns the front element of the queue.
// Returns an error if the queue is empty.
func (q *queue[T]) Pop() (T, error) {
	if q.Empty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, nil
}

// Clear removes all elements from the queue.
func (q *queue[T]) Clear() {
	q.elements = nil
}

// ToSlice returns a copy of the queue's elements as a slice.
func (q *queue[T]) ToSlice() []T {
	result := make([]T, q.Size())
	copy(result, q.elements)
	return result
}
