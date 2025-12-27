package stack

import "github.com/tshivamkr/go-dsa/linkedlist"

type StackLL[T comparable] struct {
	data *linkedlist.SinglyLL[T]
}

func NewStackLL[T comparable]() *StackLL[T] {
	return &StackLL[T]{}
}

func (s *StackLL[T]) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s *StackLL[T]) Size() int {
	return s.data.Size()
}

func (s *StackLL[T]) Push(v T) {
	s.data.Append(v)
}

func (s *StackLL[T]) Pop() (T, bool) {
	var element T
	if s.IsEmpty() {
		return element, false
	}

	// travsere to end and remove it and return the element
	deletedEl, _ := s.data.DeleteAtEnd()
	return deletedEl, true
}

func (s *StackLL[T]) Peek() (T, bool) {
	var element T
	if s.IsEmpty() {
		return element, false
	}

	element, _ = s.data.Last()
	return element, true
}

func (s *StackLL[T]) Clear() {
	s.data = nil
}
