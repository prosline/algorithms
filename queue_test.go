package main

import (
	"testing"
)

func TestQueue(t *testing.T){
	q := Queue{}
	q.Enqueue(1)
	if q.slice[0] != 1{
		t.Error("Invalid number in the queue! It expects 1 and got ", q.slice[0])
	}

	v, err := q.Dequeue()
	if err != nil{
		t.Error("Error occurred, queue size is ", v)
	}
}
