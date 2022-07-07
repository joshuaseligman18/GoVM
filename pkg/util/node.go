package util

// Struct for a node
type Node struct {
	val uint // The value of the node
	next *Node // The next node in the list
}

// Function that creates a new node
func NewNode(data uint) *Node {
	node := Node { val: data }
	return &node
}

// Gets the value of the node
func (node *Node) GetVal() uint {
	return node.val
}

// Sets the value of the node
func (node *Node) SetVal(newVal uint) {
	node.val = newVal
}

// Gets the next node
func (node *Node) GetNext() *Node {
	return node.next
}

// Sets the next node
func (node *Node) SetNext(newNext *Node) {
	node.next = newNext
}