package util

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
func (queue *Queue) Dequeue() {
	if queue.head != nil {
		queue.head = queue.head.GetNext()
	}
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

// Gets the head of the queue
func (queue *Queue) GetHead() *Node {
	return queue.head
}