package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var element T
	if s.IsEmpty() {
		return element, false
	}

	l := len(s.data)
	element = s.data[l-1]
	var zeroVal T
	s.data[l-1] = zeroVal
	s.data = s.data[:l-1]
	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var element T
	if s.IsEmpty() {
		return element, false
	}

	l := len(s.data)
	element = s.data[l-1]
	return element, true
}

func (s *Stack[T]) Clear() {
	s.data = nil
}
