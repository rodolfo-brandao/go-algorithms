package linkedlist

import "testing"

func TestPrepend(t *testing.T) {
	l := LinkedList[int]{}

	l.New()

	l.Prepend(20)
	l.Prepend(10)

	if l.Head == nil {
		t.Fatal("The head of the list should not be 'nil' after prepending an element.")
	}

	if l.Head.Value != 10 {
		t.Errorf("Expected the head of the list to be 10, but got %d", l.Head.Value)
	}

	if l.Tail.Value != 20 {
		t.Errorf("Expected the tail of the lsit to be 20, but got %d", l.Tail.Value)
	}
}

func TestAppend(t *testing.T) {
	l := LinkedList[int]{}

	l.New()

	l.Append(10)
	l.Append(20)

	if l.Head == nil {
		t.Fatal("The head of the list should not be 'nil' after appending an element.")
	}

	if l.Head.Value != 10 {
		t.Errorf("Expected the head of the list to be 10, but got %d", l.Head.Value)
	}

	if l.Tail.Value != 20 {
		t.Errorf("Expected the tail of the lsit to be 20, but got %d", l.Tail.Value)
	}
}

func TestToString(t *testing.T) {
	l := LinkedList[int]{}

	l.New()

	l.Append(10)
	l.Append(20)
	l.Append(30)
	l.Append(40)
	l.Append(50)

	result := l.String()
	expected := "10, 20, 30, 40, 50"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
