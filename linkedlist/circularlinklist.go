package linkedlist

type SCLLNode[T comparable] struct {
	Val  T
	Next *SCLLNode[T]
}

type CircularSLL[T comparable] struct {
	Head *SCLLNode[T]
	Tail *SCLLNode[T]
}

type DCLLNode[T comparable] struct {
	Val  T
	Prev *DCLLNode[T]
	Next *DCLLNode[T]
}

type CircularDLL[T comparable] struct {
	Head *DCLLNode[T]
	Tail *DCLLNode[T]
}

func NewSCLLNode[T comparable](v T) *SCLLNode[T] {
	return &SCLLNode[T]{Val: v}
}

func NewSCLL[T comparable](v T) *CircularDLL[T] {
	return &CircularDLL[T]{}
}

func (csll *CircularSLL[T]) Append(v T) bool {
	n := NewSCLLNode(v)

	if csll.Head == nil {
		n.Next = n
		csll.Head = n
		csll.Tail = n
		return true
	}

	n.Next = csll.Head
	csll.Tail.Next = n
	csll.Tail = n
	return true
}
