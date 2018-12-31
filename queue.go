package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(n int) {
	//TODO: Add element to the queue
	q.slice = append(q.slice, n)
}

func (q *Queue) Dequeue() (int, error) {
	//TODO: Extract the first element in the queue
	var err error

	if len(q.slice) > 0{
		value := q.slice[0]
		if len(q.slice) > 1{
			q.slice = q.slice[1:(len(q.slice)-1)]
		} else {
			q.slice = []int{}
		}
		return value, nil

	} else {
		err = errors.New("Queue is Empty")
		return -1, err
	}

}
func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}
