package linkedlist

type SLLNode[T comparable] struct {
	Val  T
	Next *SLLNode[T]
}

func NewNode[T comparable](v T) *SLLNode[T] {
	return &SLLNode[T]{Val: v}
}

type SinglyLL[T comparable] struct {
	size int
	head *SLLNode[T]
}

func New[T comparable]() *SinglyLL[T] {
	return &SinglyLL[T]{}
}

func (sll *SinglyLL[T]) Size() int {
	return sll.size
}

func (sll *SinglyLL[T]) IsEmpty() bool {
	return sll.size == 0
}

func (sll *SinglyLL[T]) Append(v T) {
	n := NewNode(v)

	if sll.head == nil {
		sll.head = n
		sll.size++
		return
	}

	currNode := sll.head
	for ; currNode.Next != nil; currNode = currNode.Next {
	}

	currNode.Next = n
	sll.size++
}

func (sll *SinglyLL[T]) Prepend(v T) {
	n := NewNode(v)
	n.Next = sll.head
	sll.head = n
	sll.size++
}

func (sll *SinglyLL[T]) Delete(v T) (*SLLNode[T], bool) {
	if sll.head == nil {
		return nil, false
	}

	if sll.head.Val == v {
		h := sll.head
		sll.head = sll.head.Next
		h.Next = nil
		return h, true
	}

	prevNode := sll.head
	currNode := sll.head.Next
	for currNode != nil {
		if currNode.Val == v {
			prevNode.Next = currNode.Next
			currNode.Next = nil
			return currNode, true
		}

		prevNode = currNode
		currNode = currNode.Next
	}

	return nil, false
}

func (sll *SinglyLL[T]) Find(v T) (*SLLNode[T], bool) {
	if sll.head == nil {
		return nil, false
	}

	currNode := sll.head
	for currNode != nil {
		if currNode.Val == v {
			return currNode, true
		}

		currNode = currNode.Next
	}

	return nil, false
}

func (sll *SinglyLL[T]) IsCyclicSLL() bool {
	if sll.head == nil {
		return false
	}

	slow := sll.head
	fast := sll.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
