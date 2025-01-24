package main

// Node represents a node in both queue and list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Queue represents a generic queue
type Queue[T any] struct {
	Front *Node[T]
}

// Push adds an element to the front of the queue
func (q *Queue[T]) Push(value T) {
	newNode := &Node[T]{Value: value}
	if q.Front == nil {
		q.Front = newNode
		return
	}
	current := q.Front
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Pop removes and returns the front element from the queue
// Returns zero value and false if queue is empty
func (q *Queue[T]) Pop() (T, bool) {
	if q.Front == nil {
		var zero T
		return zero, false
	}
	value := q.Front.Value
	q.Front = q.Front.Next
	return value, true
}

// List represents a generic linked list
type List[T any] struct {
	Head *Node[T]
}

// NewList returns a new empty list
func NewList[T any]() *List[T] {
	return &List[T]{
		Head: nil,
	}
}

func newNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val, Next: nil}
}

// AddNewNode adds a new node at the end of the list
func (l *List[T]) AddNewNode(val T) {
	node := newNode(val)
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}
