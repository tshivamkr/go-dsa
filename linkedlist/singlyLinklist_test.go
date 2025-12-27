package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	list := New[int]()

	if list == nil {
		t.Fatal("expected list, got nil")
	}

	if !list.IsEmpty() {
		t.Errorf("expected list to be empty")
	}

	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
}

func listToSlice[T comparable](l *SinglyLL[T]) []T {
	var result []T
	for n := l.head; n != nil; n = n.Next {
		result = append(result, n.Val)
	}
	return result
}

func TestAppend(t *testing.T) {
	list := New[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	expected := []int{10, 20, 30}
	got := listToSlice(list)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestPrepend(t *testing.T) {
	list := New[int]()

	list.Prepend(10)
	list.Prepend(20)

	expected := []int{20, 10}
	got := listToSlice(list)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestDeleteMiddle(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	_, ok := list.Delete(2)
	if !ok {
		t.Fatalf("expected delete to succeed")
	}

	expected := []int{1, 3}
	if !reflect.DeepEqual(listToSlice(list), expected) {
		t.Errorf("expected %v, got %v", expected, listToSlice(list))
	}
}

func TestDeleteHead(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)

	_, ok := list.Delete(1)
	if !ok {
		t.Fatalf("expected delete to succeed")
	}

	expected := []int{2}
	if !reflect.DeepEqual(listToSlice(list), expected) {
		t.Errorf("expected %v, got %v", expected, listToSlice(list))
	}
}

func TestDeleteTail(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	_, ok := list.Delete(3)
	if !ok {
		t.Fatalf("expected delete to succeed")
	}

	expected := []int{1, 2}
	if !reflect.DeepEqual(listToSlice(list), expected) {
		t.Errorf("expected %v, got %v", expected, listToSlice(list))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	list := New[int]()
	list.Append(1)

	_, ok := list.Delete(999)
	if ok {
		t.Errorf("expected delete to fail")
	}
}

func TestFind(t *testing.T) {
	list := New[int]()
	list.Append(10)
	list.Append(20)

	node, ok := list.Find(20)
	if !ok {
		t.Fatalf("expected to find value")
	}

	if node.Val != 20 {
		t.Errorf("expected 20, got %d", node.Val)
	}
}

func TestFindNotFound(t *testing.T) {
	list := New[int]()
	list.Append(1)

	_, ok := list.Find(999)
	if ok {
		t.Errorf("expected not found")
	}
}

func TestSize(t *testing.T) {
	list := New[int]()

	if list.Size() != 0 {
		t.Errorf("expected size 0")
	}

	list.Append(1)
	list.Append(2)

	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	list := New[int]()

	if !list.IsEmpty() {
		t.Errorf("expected list to be empty")
	}

	list.Append(1)

	if list.IsEmpty() {
		t.Errorf("expected list to be non-empty")
	}
}

func TestMultipleOperations(t *testing.T) {
	list := New[int]()

	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	list.Delete(1)
	list.Append(3)

	expected := []int{0, 2, 3}

	if !reflect.DeepEqual(listToSlice(list), expected) {
		t.Errorf("expected %v, got %v", expected, listToSlice(list))
	}
}

func buildList[T comparable](vals ...T) *SinglyLL[T] {
	if len(vals) == 0 {
		return &SinglyLL[T]{}
	}

	head := &SLLNode[T]{Val: vals[0]}
	curr := head

	for _, v := range vals[1:] {
		curr.Next = &SLLNode[T]{Val: v}
		curr = curr.Next
	}

	return &SinglyLL[T]{head: head}
}

// create cycle: last node points to node at index pos
func makeCycle[T comparable](sll *SinglyLL[T], pos int) {
	if sll.head == nil {
		return
	}

	var target *SLLNode[T]
	curr := sll.head
	index := 0

	for curr.Next != nil {
		if index == pos {
			target = curr
		}
		curr = curr.Next
		index++
	}

	if target != nil {
		curr.Next = target
	}
}

//
// ================= TESTS =================
//

func TestIsCyclic_EmptyList(t *testing.T) {
	sll := &SinglyLL[int]{}

	if sll.IsCyclicSLL() {
		t.Fatal("expected no cycle in empty list")
	}
}

func TestIsCyclic_SingleNode_NoCycle(t *testing.T) {
	sll := buildList(1)

	if sll.IsCyclicSLL() {
		t.Fatal("expected no cycle")
	}
}

func TestIsCyclic_SingleNode_SelfCycle(t *testing.T) {
	sll := buildList(1)
	sll.head.Next = sll.head

	if !sll.IsCyclicSLL() {
		t.Fatal("expected cycle")
	}
}

func TestIsCyclic_MultipleNodes_NoCycle(t *testing.T) {
	sll := buildList(1, 2, 3, 4)

	if sll.IsCyclicSLL() {
		t.Fatal("expected no cycle")
	}
}

func TestIsCyclic_CycleAtHead(t *testing.T) {
	sll := buildList(1, 2, 3, 4)
	makeCycle(sll, 0)

	if !sll.IsCyclicSLL() {
		t.Fatal("expected cycle at head")
	}
}

func TestIsCyclic_CycleInMiddle(t *testing.T) {
	sll := buildList(1, 2, 3, 4, 5)
	makeCycle(sll, 2)

	if !sll.IsCyclicSLL() {
		t.Fatal("expected cycle in middle")
	}
}

func TestIsCyclic_TailSelfLoop(t *testing.T) {
	sll := buildList(1, 2, 3)

	// create self-loop at tail
	curr := sll.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = curr

	if !sll.IsCyclicSLL() {
		t.Fatal("expected self-loop cycle")
	}
}
