# Go Data Structures – `dsgo`

A generic data structures library in Go, designed to be simple, reusable, and efficient.

---

## 📦 Installation

To use this repository in your Go project, run:

```go
go get github.com/fuad7161/dsgo
```
## Jump to Specific Data Structures

| Linear Data Structures            | Tree-Based Structures         | Hash-Based Structures     | Heap Structures           | Graph Structures              | Others / Advanced                   |
|----------------------------------|------------------------------|---------------------------|---------------------------|-------------------------------|-----------------------------------|
| [Stack](#stack)                  | [Binary Tree](#binary-tree)  | [Hash Table](#hash-table) | [Min Heap](#min-heap)     | [Graph](#graph)               | [LRU Cache](#lru-cache)            |
| [Queue](#queue)                  | [Binary Search Tree](#binary-search-tree) | [Bloom Filter](#bloom-filter) | [Max Heap](#max-heap)       | [Union-Find (Disjoint Set)](#union-find-disjoint-set) | [Skip List](#skip-list)            |
| [Circular Queue](#circular-queue) | [AVL Tree](#avl-tree)         |                           | [D-ary Heap](#d-ary-heap) |                               | [Suffix Array](#suffix-array)      |
| [Priority Queue](#priority-queue) | [Red-Black Tree](#red-black-tree) |                           | [Fibonacci Heap](#fibonacci-heap) |                           | [Suffix Tree](#suffix-tree)        |
| [Deque](#deque)                  | [Segment Tree](#segment-tree) |                           |                           |                               | [Fenwick Tree (Binary Indexed Tree)](#fenwick-tree-binary-indexed-tree) |
| [Singly Linked List](#singly-linked-list) | [Trie (Prefix Tree)](#trie-prefix-tree) |                           |                           |                               | [KD Tree](#kd-tree)                |
| [Doubly Linked List](#doubly-linked-list) | [B-Tree](#b-tree)             |                           |                           |                               |                                   |
| [Circular Linked List](#circular-linked-list) | [B+ Tree](#b-tree-1)          |                           |                           |                               |                                   |

## Stack
🧭 Stack Methods Overview |  [🔝 Back to Top](#Jump-to-Specific-Data-Structures)
```go
func main() {
	// Create a stack of integers
	stack := dsgo.NewStack[int]()

	// Push elements
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Peek top element
	top, _ := stack.Peek()
	fmt.Println("Top:", top) // Output: Top: 30

	// Pop elements
	val, _ := stack.Pop()
	fmt.Println("Popped:", val) // Output: Popped: 30

	// Stack length
	fmt.Println("Length:", stack.Len()) // Output: Length: 2

	// Check if empty
	fmt.Println("Is empty?", stack.Empty()) // Output: Is empty? false

	// Convert to slice
	slice := stack.ToSlice()
	fmt.Println("Slice:", slice) // Output: Slice: [10 20]

	// Clear the stack
	stack.Clear()
	fmt.Println("Cleared. Is empty?", stack.Empty()) // Output: Cleared. Is empty? true
}

```

## Queue
🧭 Stack Methods Overview |  [🔝 Back to Top](#Jump-to-Specific-Data-Structures)
```go
func main() {
	// Create a queue of integers
	queue := dsgo.NewQueue[int]()

	// Push elements
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)

	// Peak front element
	front, _ := queue.Peak()
	fmt.Println("Front:", front) // Output: Front: 10

	// Pop elements
	val, _ := queue.Pop()
	fmt.Println("Popped:", val) // Output: Popped: 10

	// Queue length
	fmt.Println("Length:", queue.Len()) // Output: Length: 2

	// Check if empty
	fmt.Println("Is empty?", queue.Empty()) // Output: Is empty? false

	// Convert to slice
	slice := queue.ToSlice()
	fmt.Println("Slice:", slice) // Output: Slice: [20 30]

	// Clear the queue
	queue.Clear()
	fmt.Println("Cleared. Is empty?", queue.Empty()) // Output: Cleared. Is empty? true
}
```
