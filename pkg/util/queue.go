package util

import (
	"fmt"
	"strings"
)

// Struct for a queue
type Queue struct {
	head *Node // The head of the queue
}

// Creates a new empty queue
func NewQueue() *Queue {
	queue := Queue { head: nil }
	return &queue
}

// Adds the value to the end of the queue
func (queue *Queue) Enqueue(val uint32) {
	newNode := NewNode(val)
	if queue.head == nil {
		queue.head = newNode
	} else {
		cur := queue.head
		for cur.GetNext() != nil {
			cur = cur.GetNext()
		}
		cur.SetNext(newNode)
	}
}

// Removes the top element from the queue
func (queue *Queue) Dequeue() uint32 {
	if queue.head != nil {
		val := queue.head.GetVal()
		queue.head = queue.head.GetNext()
		return val
	}
	return 0
}

// Function that determines if a given value is in the queue
func (queue *Queue) Contains(val uint32) bool {
	cur := queue.head
	for cur != nil {
		if cur.GetVal() == val {
			return true
		} else {
			cur = cur.GetNext()
		}
	}
	return false
}

// Function that removes the last item in the queue
func (queue *Queue) RemoveLast() {
	cur := queue.head
	if cur == nil || cur.next == nil {
		queue.head = nil
	}
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
}

// Gets the head of the queue
func (queue *Queue) GetHead() *Node {
	return queue.head
}

// Gets the string representation of the queue
func (queue *Queue) ToString() string {
	var str strings.Builder
	cur := queue.head

	for cur != nil {
		str.WriteString(fmt.Sprintf("%d ", cur.val))
		cur = cur.next
	}
	return str.String()
}