package main

import (
	"errors"
	"fmt"
	"sort"
)

type Stack struct {
	slice []int
}

func (s *Stack) getMin() int {
	tmpSlice := s.slice
	sort.Ints(tmpSlice)
	fmt.Println(tmpSlice)
	return tmpSlice[0]
}

func (s *Stack) Push(n int) {
	//TODO: Add element to the stack
	s.slice = append(s.slice, n)
}

func (s *Stack) Peek() (int, error) {
	var value int
	var err error
	if len(s.slice) > 0 {
		value = s.slice[len(s.slice)-1]
	} else {
		err = errors.New("Stack is empty")
		return 0, err
	}
	return value, nil
}
func (s *Stack) Pop() (int, error) {
	//TODO: Extract the last element of the stack
	var err error
	var ret int

	ret, err = s.Peek()
	if err != nil {
		err = errors.New("Stack is Empty")
		return -1, err
	}

	s.slice = s.slice[0:(len(s.slice) - 2)]
	return ret, nil
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}
