package dsgo

import "fmt"

// Queue is a generic queue implementation using a slice.
type queue[T any] struct {
	elements []T
}

// NewQueue creates and returns a new empty Queue.
func NewQueue[T any]() *queue[T] {
	return &queue[T]{make([]T, 0)}
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

// Peek returns the front element without removing it.
// Returns an error if the queue is empty.
func (q *queue[T]) Peek() (T, error) {
	if q.Empty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

// Len returns the number of elements in the queue.
func (q *queue[T]) Len() int {
	return len(q.elements)
}

// Empty checks if the queue has no elements.
func (q *queue[T]) Empty() bool {
	return q.Len() == 0
}

// Clear removes all elements from the queue.
func (q *queue[T]) Clear() {
	q.elements = nil
}

// ToSlice returns a copy of the queue's elements as a slice.
func (q *queue[T]) ToSlice() []T {
	result := make([]T, q.Len())
	copy(result, q.elements)
	return result
}
