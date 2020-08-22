package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(-1)
	value, _ := s.Peek()

	if value != -1 {
		t.Error("Invalid element returned by Stack! It expects 2 and got ", value)
	}

	v, err := s.Pop()
	if err != nil {
		t.Error("Error occurred, Stack size is ", v)
	}
	n := s.getMin()
	if n != 1 {
		t.Error("Error occurred, minimun number should be 1 instead is  ", 1)
	}
}
