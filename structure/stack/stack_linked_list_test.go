package stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	t.Log(s)
}

func TestLinkedListStack_Pop(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 3; i++ {
		s.Push(i)
	}

	if s.Pop() != 2 {
		t.Fatal("pop failed")
	}
	if s.Pop() != 1 {
		t.Fatal("pop failed")
	}
	if s.Pop() != 0 {
		t.Fatal("pop failed")
	}
	if s.Pop() != nil {
		t.Fatal("pop failed")
	}
	t.Log(s)
}

func TestLinkedListStack_Top(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)

	if s.Top() != 1 {
		t.Fatal("top failed")
	}
	if s.Pop(); s.Top() != nil {
		t.Fatal("top empty stack failed")
	}
}

func TestLinkedListStack_Flush(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.Flush(); !s.IsEmpty() {
		t.Fatal("flush failed")
	}
}
