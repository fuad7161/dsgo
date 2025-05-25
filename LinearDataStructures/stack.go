package dsgo

import "fmt"

// Stack is a generic stack implementation using a slice.
type stack[T any] struct {
	elements []T
}

func StackInt() *stack[int] {
	return &stack[int]{make([]int, 0)}
}

func StackInt8() *stack[int8] {
	return &stack[int8]{make([]int8, 0)}
}

func StackInt16() *stack[int16] {
	return &stack[int16]{make([]int16, 0)}
}

func StackInt32() *stack[int32] {
	return &stack[int32]{make([]int32, 0)}
}

func StackInt64() *stack[int64] {
	return &stack[int64]{make([]int64, 0)}
}

func StackUint() *stack[uint] {
	return &stack[uint]{make([]uint, 0)}
}

func StackUint8() *stack[uint8] {
	return &stack[uint8]{make([]uint8, 0)}
}

func StackUint16() *stack[uint16] {
	return &stack[uint16]{make([]uint16, 0)}
}

func StackUint32() *stack[uint32] {
	return &stack[uint32]{make([]uint32, 0)}
}

func StackUint64() *stack[uint64] {
	return &stack[uint64]{make([]uint64, 0)}
}

func StackFloat32() *stack[float32] {
	return &stack[float32]{make([]float32, 0)}
}

func StackFloat64() *stack[float64] {
	return &stack[float64]{make([]float64, 0)}
}

func StackComplex64() *stack[complex64] {
	return &stack[complex64]{make([]complex64, 0)}
}

func StackComplex128() *stack[complex128] {
	return &stack[complex128]{make([]complex128, 0)}
}

func StackString() *stack[string] {
	return &stack[string]{make([]string, 0)}
}

func StackBool() *stack[bool] {
	return &stack[bool]{make([]bool, 0)}
}

func StackByte() *stack[byte] {
	return &stack[byte]{make([]byte, 0)}
}

func StackRune() *stack[rune] {
	return &stack[rune]{make([]rune, 0)}
}

// NewStack creates and returns a new empty Stack.
func NewStack[T any]() *stack[T] {
	return &stack[T]{make([]T, 0)}
}

// Empty checks whether the stack has no elements.
func (s *stack[T]) Empty() bool {
	return s.Size() == 0
}

// Peek returns the top element without removing it from the stack.
// Returns an error if the stack is empty.
func (s *stack[T]) Peek() (T, error) {
	return s.Back()
}

// Size returns the number of elements in the stack.
func (s *stack[T]) Size() int {
	return len(s.elements)
}

// Push adds an element to the top of the stack.
func (s *stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

// Pop removes and returns the top element of the stack.
// Returns an error if the stack is empty.
func (s *stack[T]) Pop() (T, error) {
	if s.Empty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	top, _ := s.Back()
	s.elements = s.elements[:s.Size()-1]
	return top, nil
}

// Back returns the top element of the stack without removing it.
// Returns an error if the stack is empty.
func (s *stack[T]) Back() (T, error) {
	if s.Empty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.elements[s.Size()-1], nil
}

// Clear removes all elements from the stack.
func (s *stack[T]) Clear() {
	s.elements = nil
}

// ToSlice returns a copy of the underlying slice (LIFO order preserved).
func (s *stack[T]) ToSlice() []T {
	result := make([]T, s.Size())
	copy(result, s.elements)
	return result
}
