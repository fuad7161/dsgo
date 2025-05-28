package dsgo

import (
	"fmt"
)

// deque is a generic deque implementation using a slice.
type deque[T any] struct {
	elements []T
}

func DequeInt() *deque[int] {
	return &deque[int]{elements: make([]int, 0)}
}

func DequeInt8() *deque[int8] {
	return &deque[int8]{elements: make([]int8, 0)}
}

func DequeInt16() *deque[int16] {
	return &deque[int16]{elements: make([]int16, 0)}
}

func DequeInt32() *deque[int32] {
	return &deque[int32]{elements: make([]int32, 0)}
}

func DequeInt64() *deque[int64] {
	return &deque[int64]{elements: make([]int64, 0)}
}

func DequeUint() *deque[uint] {
	return &deque[uint]{elements: make([]uint, 0)}
}

func DequeUint8() *deque[uint8] {
	return &deque[uint8]{elements: make([]uint8, 0)}
}

func DequeUint16() *deque[uint16] {
	return &deque[uint16]{elements: make([]uint16, 0)}
}

func DequeUint32() *deque[uint32] {
	return &deque[uint32]{elements: make([]uint32, 0)}
}

func DequeUint64() *deque[uint64] {
	return &deque[uint64]{elements: make([]uint64, 0)}
}

func DequeFloat32() *deque[float32] {
	return &deque[float32]{elements: make([]float32, 0)}
}

func DequeFloat64() *deque[float64] {
	return &deque[float64]{elements: make([]float64, 0)}
}

func DequeComplex64() *deque[complex64] {
	return &deque[complex64]{elements: make([]complex64, 0)}
}

func DequeComplex128() *deque[complex128] {
	return &deque[complex128]{elements: make([]complex128, 0)}
}

func DequeString() *deque[string] {
	return &deque[string]{elements: make([]string, 0)}
}

func DequeBool() *deque[bool] {
	return &deque[bool]{elements: make([]bool, 0)}
}

func DequeByte() *deque[byte] {
	return &deque[byte]{elements: make([]byte, 0)}
}

func DequeRune() *deque[rune] {
	return &deque[rune]{elements: make([]rune, 0)}
}

// NewDeque creates and returns a new empty deque.
func NewDeque[T any]() *deque[T] {
	return &deque[T]{make([]T, 0)}
}

func (dq *deque[T]) PushBack(v T) {
	dq.elements = append(dq.elements, v)
}

func (dq *deque[T]) PushFront(v T) {
	dq.elements = append([]T{v}, dq.elements...)
}

func (dq *deque[T]) PopBack() (T, error) {
	if dq.Empty() {
		var v T
		return v, fmt.Errorf("deque is empty")
	}
	last, _ := dq.Back()
	dq.elements = dq.elements[:dq.Size()-1]
	return last, nil
}

func (dq *deque[T]) PopFront() (T, error) {
	if dq.Empty() {
		var v T
		return v, fmt.Errorf("deque is empty")
	}
	first, _ := dq.Front()
	dq.elements = dq.elements[1:]
	return first, nil
}

func (dq *deque[T]) Insert(idx int, v T) error {
	return nil
}

func (dq *deque[T]) Back() (T, error) {
	if dq.Empty() {
		var v T
		return v, fmt.Errorf("deque is empty")
	}
	return dq.elements[dq.Size()-1], nil
}

func (dq *deque[T]) Front() (T, error) {
	if dq.Empty() {
		var v T
		return v, fmt.Errorf("deque is empty")
	}
	return dq.elements[0], nil
}

func (dq *deque[T]) Size() int {
	return len(dq.elements)
}
func (dq *deque[T]) Empty() bool {
	return dq.Size() == 0
}
