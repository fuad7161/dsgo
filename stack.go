package dsgo

type stack struct {
	arr []int
}

// NewStack Create a new stack
func NewStack() stack {
	return stack{}
}

// Push element into the stack
func (st *stack) Push(val int) {
	st.arr = append(st.arr, val)
}

// Pop element from the stack and remove element
func (st *stack) Pop() int {
	if st.Len() == 0 {
		return -1
	}
	var lastValue = st.arr[st.Len()-1]
	st.arr = st.arr[:st.Len()-1]
	return lastValue
}

// Peek last added element without removing
func (st *stack) Peek() int {
	if st.Len() == 0 {
		return -1
	}
	return st.arr[st.Len()-1]
}

func (st *stack) isEmpty() bool {
	return st.Len() == 0
}

func (st *stack) Len() int {
	return len(st.arr)
}
