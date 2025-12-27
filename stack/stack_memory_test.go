package stack

import "testing"

// Helper struct to track references
type bigObject struct {
	data [1024]byte
}

// ----------------------------------------------------
// Test: Pop clears reference (prevents memory retention)
// ----------------------------------------------------
func TestPopClearsReference(t *testing.T) {
	s := NewStack[*bigObject]()

	obj := &bigObject{}
	s.Push(obj)

	// Sanity check
	if s.Size() != 1 {
		t.Fatalf("expected size 1, got %d", s.Size())
	}

	// Pop element
	val, ok := s.Pop()
	if !ok {
		t.Fatal("expected pop to succeed")
	}

	if val != obj {
		t.Fatal("popped value mismatch")
	}

	// Now check internal memory
	// The slice should be empty
	if len(s.data) != 0 {
		t.Fatalf("expected empty slice, got length %d", len(s.data))
	}

	// But capacity may still exist — check backing array
	if cap(s.data) > 0 {
		// Re-slice to inspect underlying array
		internal := s.data[:cap(s.data)]

		for i, v := range internal {
			if v != nil {
				t.Fatalf("memory leak: index %d still holds reference", i)
			}
		}
	}
}
