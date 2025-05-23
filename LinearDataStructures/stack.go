package dsgo

import "fmt"

// Stack is a generic stack implementation using a slice.
type Stack[T any] struct {
	elements []T
}

// NewStack creates and returns a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

// Peek returns the top element without removing it from the stack.
// Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	return s.Back()
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

// Pop removes and returns the top element of the stack.
// Returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	top, _ := s.Back()
	s.elements = s.elements[:s.Len()-1]
	return top, nil
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

// Back returns the top element of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack[T]) Back() (T, error) {
	if s.Empty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.elements[s.Len()-1], nil
}

// Empty checks whether the stack has no elements.
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.elements = nil
}

// ToSlice returns a copy of the underlying slice (LIFO order preserved).
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, s.Len())
	copy(result, s.elements)
	return result
}
