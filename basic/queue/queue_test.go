package main

import (
	"testing"
)

func TestQueue_Pop(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	head := q.Pop()
	if head != 1 {
		t.Errorf("Pop data is %d \n", head)
	}
	if q.IsEmpty() {
		t.Fatalf("queue is empty")
	}
}
