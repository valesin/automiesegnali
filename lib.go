package main

// Node represents a node in both queue and list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// container holds minimal push/pop operations that both the queue and list can reuse.
type container[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// aggiungiInCoda adds a new element at the tail.
func (c *container[T]) aggiungiInCoda(value T) {
	newNode := &Node[T]{Value: value}
	if c.tail == nil {
		// container is empty, so head and tail both point to the new node.
		c.head = newNode
		c.tail = newNode
	} else {
		// Insert at the tail in O(1).
		c.tail.Next = newNode
		c.tail = newNode
	}
}

// cancellaInTesta removes an element from the head.
func (c *container[T]) cancellaInTesta() (T, bool) {
	if c.head == nil {
		var zero T
		return zero, false // Empty container
	}
	val := c.head.Value
	c.head = c.head.Next
	if c.head == nil {
		// If the container becomes empty, tail should also be nil.
		c.tail = nil
	}
	return val, true
}

// Queue embeds container and provides Enqueue/Dequeue methods.
type Queue[T any] struct {
	container[T]
}

// Enqueue is a public-facing method that calls the shared push logic.
func (q *Queue[T]) Enqueue(value T) {
	q.aggiungiInCoda(value)
}

// Dequeue is a public-facing method that calls the shared pop logic.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.cancellaInTesta()
}

// LinkedList embeds container for Add/Remove methods (or other list-specific logic).
type LinkedList[T any] struct {
	container[T]
}

// Add adds a node at the end, but you could easily change it to insert at the front.
func (ll *LinkedList[T]) Add(value T) {
	ll.aggiungiInCoda(value)
}

// Below is one possible approach to reduce code repetition while still exposing both a “linked list” interface and a “queue” interface. You can define a shared container that handles the low-level node operations, then embed that container in your specialized types. This way, the underlying logic lives in just one place, while the separate types each provide clarity to anyone reading or using the code.
//Explanation:
//• container[T] centralizes common operations (push/pop).
//• Queue[T] and LinkedList[T] embed container[T], directly reusing the logic without additional code.
//• Each specialized type uses different method names (like Enqueue/Dequeue vs. Add/Remove) to clarify intent.
// Decido di inserire in coda e togliere dalla testa, sia per mantenere coerenza con l'ordine di inserimento che per evitare di dover scorrere tutta la lista per inserire in coda.
// La roba che faccio si chiama embedding, e mi permette di riutilizzare codice senza doverlo riscrivere
// This works because of Go's struct embedding feature. When LinkedList embeds container, all the fields and methods of container are automatically promoted to LinkedList. This means:

//LinkedList inherits all fields from container
//You can access head directly through LinkedList without mentioning container
//Go treats it as if head was directly declared in LinkedList
// Quindi posso accedere a head direttamente da LinkedList senza dover menzionare container
