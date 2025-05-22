# Repo for Go Data Structure
### To use this repo just get this into go project
```go
go get github.com/fuad7161/dsgo
```
## Heap
### Key functionality for Heap
```go
func main() {

    hp := dsgo.NewHeap() // Create a new Heap (Min heap)
    hp.Insert(100)       // Insert value into the heap
    hp.Pop()             // Pop the minimum element from the heap
}
```

## Stack
```go
func main() {

	st := dsgo.NewStack()

	st.Push(100) // Push value into stack
	st.Len()     // Get the length of the stack
	st.Peek()    // Get the last added element but not removing from stack
	st.Pop()     // Get the last added element and removed from stack
	st.IsEmpty() // Stack is Empty or not checking
}
```
