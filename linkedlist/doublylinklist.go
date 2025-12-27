package linkedlist

type DLLNode[T comparable] struct {
	Val  T
	Prev *DLLNode[T]
	Next *DLLNode[T]
}

type DoublyLL[T comparable] struct {
	length int
	head   *DLLNode[T]
}

func NewDLLNode[T comparable](v T) *DLLNode[T] {
	return &DLLNode[T]{Val: v}
}

func NewDLL[T comparable]() *DoublyLL[T] {
	return &DoublyLL[T]{}
}

func (dll *DoublyLL[T]) Append(v T) {
	n := NewDLLNode(v)

	if dll.head == nil {
		dll.head = n
		dll.length++
		dll.head.Prev = nil
		dll.head.Next = nil
		return
	}

	currNode := dll.head
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = n
	n.Prev = currNode
}

func (dll *DoublyLL[T]) Delete(v T) (*DLLNode[T], bool) {
	if dll.head == nil {
		return nil, false
	}

	if dll.head.Val == v {
		deleted := dll.head
		dll.head = dll.head.Next
		if dll.head != nil {
			dll.head.Prev = nil
		}
		deleted.Next = nil
		deleted.Prev = nil
		return deleted, true
	}

	currNode := dll.head.Next
	for currNode != nil {
		if currNode.Val == v {
			currNode.Prev.Next = currNode.Next
			if currNode.Next != nil {
				currNode.Next.Prev = currNode.Prev
			}
			currNode.Prev = nil
			currNode.Next = nil
			return currNode, true
		}
		currNode = currNode.Next
	}

	return nil, false
}

func (dll *DoublyLL[T]) IsCyclicSLL() bool {
	if dll.head == nil {
		return false
	}

	slow := dll.head
	fast := dll.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func (dll DoublyLL[T]) HasBrokenPrevLinks() bool {
	var prev *DLLNode[T]
	curr := dll.head

	for curr != nil {
		if curr.Prev != prev {
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}
