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
func (queue *Queue) Enqueue(val uint) {
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