package linkedlist

import (
	"reflect"
	"testing"
)

//
// ===== Helpers =====
//

func listToSliceForward[T comparable](dll *DoublyLL[T]) []T {
	var res []T
	for curr := dll.head; curr != nil; curr = curr.Next {
		res = append(res, curr.Val)
	}
	return res
}

// Get tail by walking forward
func getTail[T comparable](dll *DoublyLL[T]) *DLLNode[T] {
	curr := dll.head
	if curr == nil {
		return nil
	}
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}

// Backward traversal starting from tail
func listToSliceBackward[T comparable](dll *DoublyLL[T]) []T {
	var res []T
	curr := getTail(dll)
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Prev
	}
	return res
}

//
// ===== TESTS =====
//

func TestDoublyLL_Append(t *testing.T) {
	ll := &DoublyLL[int]{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	want := []int{1, 2, 3}
	got := listToSliceForward(ll)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	back := listToSliceBackward(ll)
	wantBack := []int{3, 2, 1}

	if !reflect.DeepEqual(back, wantBack) {
		t.Fatalf("reverse expected %v, got %v", wantBack, back)
	}
}

func TestDelete_EmptyList(t *testing.T) {
	ll := &DoublyLL[int]{}

	node, ok := ll.Delete(10)
	if ok || node != nil {
		t.Fatal("expected delete to fail on empty list")
	}
}

func TestDelete_SingleNode(t *testing.T) {
	ll := &DoublyLL[int]{}
	ll.Append(10)

	node, ok := ll.Delete(10)
	if !ok || node.Val != 10 {
		t.Fatal("failed to delete existing node")
	}

	if ll.head != nil {
		t.Fatal("list should be empty after deletion")
	}
}

func TestDelete_Head(t *testing.T) {
	ll := &DoublyLL[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Delete(1)

	want := []int{2, 3}
	got := listToSliceForward(ll)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	if ll.head.Prev != nil {
		t.Fatal("head.prev should be nil")
	}
}

func TestDelete_Tail(t *testing.T) {
	ll := &DoublyLL[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Delete(3)

	want := []int{1, 2}
	got := listToSliceForward(ll)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestDelete_Middle(t *testing.T) {
	ll := &DoublyLL[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Delete(2)

	want := []int{1, 3}
	got := listToSliceForward(ll)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	if ll.head.Next.Prev != ll.head {
		t.Fatal("broken Prev pointer")
	}
}

func TestDelete_NotFound(t *testing.T) {
	ll := &DoublyLL[int]{}
	ll.Append(1)
	ll.Append(2)

	node, ok := ll.Delete(99)

	if ok || node != nil {
		t.Fatal("expected delete to fail")
	}

	want := []int{1, 2}
	got := listToSliceForward(ll)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("list mutated unexpectedly")
	}
}

// func TestDoublyLL_NoCycles(t *testing.T) {
// 	ll := &DoublyLL[int]{}
// 	ll.Append(1)
// 	ll.Append(2)
// 	ll.Append(3)

// 	visited := map[*DLLNode[int]]bool{}
// 	curr := ll.head

// 	for curr != nil {
// 		if visited[curr] {
// 			t.Fatal("cycle detected")
// 		}
// 		visited[curr] = true
// 		curr = curr.Next
// 	}
// }
