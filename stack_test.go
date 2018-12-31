package main

import (
	"testing"
)

func TestStack(t *testing.T){
	s := Stack{}
	s.Push(1)
	s.Push(2)
	//s.Push(3)
	value, err := s.Peek()

	if value != 2{
		t.Error("Invalid element returned by Stack! It expects 2 and got ", value)
	}

	v, err := s.Pop()
	if err != nil {
		t.Error("Error occurred, Stack size is ", v)
	}
}

