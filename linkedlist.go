package main

import "fmt"

type LinkNode struct {
	prev *LinkNode
	next *LinkNode
	key interface{}
}

type List struct {
	head *LinkNode
	tail *LinkNode
	length int
}

func (l *List) Insert(i interface{}){
	list := &LinkNode{
		next: l.head,
		key: i,
	}
	if l.head != nil {
		l.head.prev = list
	}
	l.head = list

	lst := l.head
	for lst.next != nil{
		lst = lst.next
	}
	l.tail = lst
	l.length++
}

func (l *List) Delete(value int){
	if l.length == 0{
		return
	}
	valueDelete := l.head
	for valueDelete.next.key != value {
		valueDelete = valueDelete.next
	}
	valueDelete.next = valueDelete.next.next

}
func (l *List) Display(){
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}
func Display(l *LinkNode){
	for l != nil {
		fmt.Printf("%v -> ",l.key)
		l = l.next
	}
	fmt.Println()
}

func (l *List) Reverse(){
	current := l.head
	var prev *LinkNode
	l.tail = l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	Display(l.head)
}

func main() {

	l := List{}
	l.Insert(5)
	l.Insert(6)
	l.Insert(8)
	l.Insert(10)
	l.Insert(12)
	l.Insert(11)
	l.Insert(0)
	fmt.Println("\n================================\n")
	fmt.Printf("Head: %v\n", l.head.key)
	fmt.Printf("Tail: %v\n", l.tail.key)
	l.Display()
	fmt.Println("\n================================\n")
	fmt.Printf("Head: %v\n", l.head.key)
	fmt.Printf("Tail: %v\n", l.tail.key)
	l.Reverse()
	fmt.Println("\n================================\n")
	l.Delete(10)
	fmt.Println("\n================================\n")
	l.Display()
}
