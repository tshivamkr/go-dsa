package stack

import "testing"

// ------------------------------
// Test: New Stack
// ------------------------------
func TestNewStack(t *testing.T) {
	s := NewStack[int]()

	if s == nil {
		t.Fatal("expected stack to be initialized, got nil")
	}

	if !s.IsEmpty() {
		t.Errorf("expected new stack to be empty")
	}

	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}
}

// ------------------------------
// Test: Push
// ------------------------------
func TestPush(t *testing.T) {
	s := NewStack[int]()

	s.Push(10)
	s.Push(20)

	if s.Size() != 2 {
		t.Fatalf("expected size 2, got %d", s.Size())
	}

	top, ok := s.Peek()
	if !ok || top != 20 {
		t.Fatalf("expected top 20, got %v", top)
	}
}

// ------------------------------
// Test: Pop
// ------------------------------
func TestPop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, ok := s.Pop()
	if !ok || val != 3 {
		t.Fatalf("expected 3, got %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 2 {
		t.Fatalf("expected 2, got %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %v", val)
	}

	if !s.IsEmpty() {
		t.Fatal("stack should be empty after popping all elements")
	}
}

// ------------------------------
// Test: Pop on Empty Stack
// ------------------------------
func TestPopEmpty(t *testing.T) {
	s := NewStack[int]()

	val, ok := s.Pop()
	if ok {
		t.Fatalf("expected pop to fail, got %v", val)
	}
}

// ------------------------------
// Test: Peek
// ------------------------------
func TestPeek(t *testing.T) {
	s := NewStack[int]()

	if _, ok := s.Peek(); ok {
		t.Fatal("peek should fail on empty stack")
	}

	s.Push(42)

	val, ok := s.Peek()
	if !ok || val != 42 {
		t.Fatalf("expected 42, got %v", val)
	}

	// Ensure peek does not remove element
	if s.Size() != 1 {
		t.Fatalf("peek should not remove element")
	}
}

// ------------------------------
// Test: LIFO Order
// ------------------------------
func TestLIFOOrder(t *testing.T) {
	s := NewStack[int]()

	values := []int{10, 20, 30, 40}
	for _, v := range values {
		s.Push(v)
	}

	for i := len(values) - 1; i >= 0; i-- {
		val, ok := s.Pop()
		if !ok || val != values[i] {
			t.Fatalf("expected %d, got %d", values[i], val)
		}
	}
}

// ------------------------------
// Test: Large Volume (Stress Test)
// ------------------------------
func TestLargeStack(t *testing.T) {
	s := NewStack[int]()
	const n = 100_000

	for i := 0; i < n; i++ {
		s.Push(i)
	}

	if s.Size() != n {
		t.Fatalf("expected size %d, got %d", n, s.Size())
	}

	for i := n - 1; i >= 0; i-- {
		val, ok := s.Pop()
		if !ok || val != i {
			t.Fatalf("expected %d, got %d", i, val)
		}
	}
}

// ------------------------------
// Table-Driven Test
// ------------------------------
func TestStack_TableDriven(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"single", []int{1}, []int{1}},
		{"multiple", []int{1, 2, 3}, []int{3, 2, 1}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()

			for _, v := range tt.input {
				s.Push(v)
			}

			for _, expected := range tt.expect {
				val, ok := s.Pop()
				if !ok || val != expected {
					t.Fatalf("expected %d, got %d", expected, val)
				}
			}

			if !s.IsEmpty() {
				t.Fatalf("stack should be empty after test")
			}
		})
	}
}
