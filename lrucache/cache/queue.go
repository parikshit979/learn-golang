package cache

import (
	"log"
)

const (
	// DefaultSize is the default size of the cache.
	DefaultSize = 100
)

// Node is a doubly linked list node.
type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

// Queue is a doubly linked list.
type Queue struct {
	head     *Node
	tail     *Node
	size     int
	capacity int
}

// NewQueue creates a new Queue.
func NewQueue(capacity int) *Queue {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &Queue{
		head:     head,
		tail:     tail,
		capacity: capacity,
		size:     0,
	}
}

// Add adds a node to the front of the queue.
func (q *Queue) Add(key, value string) {
	log.Printf("Adding node to queue: key(%s), value(%s)", key, value)
	// If the queue is full, remove the last node.
	if q.size >= q.capacity {
		tempnode := q.Pop()
		log.Printf("Queue is full, removing last node: key(%s), value(%s)",
			tempnode.key, tempnode.value)
		q.size--
	}

	node := &Node{
		key:   key,
		value: value,
		prev:  q.head,
		next:  q.head.next,
	}
	q.head.next.prev = node
	q.head.next = node
	q.size++
}

// Remove removes a node from the queue.
func (q *Queue) Remove(node *Node) {
	log.Printf("Removing node from queue: key(%s)", node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
	q.size--
}

// MoveToFront moves a node to the front of the queue.
func (q *Queue) MoveToFront(node *Node) {
	log.Printf("Moving node to front of queue: key(%s)", node.key)
	q.Remove(node)
	q.Add(node.key, node.value)
}

// Pop removes the last node from the queue.
func (q *Queue) Pop() *Node {
	node := q.tail.prev
	log.Printf("Popping node from queue: key(%s), value(%s)", node.key, node.value)
	q.Remove(node)
	return node
}
