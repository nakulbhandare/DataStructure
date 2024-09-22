package main

import "fmt"

type Node struct {
	key  int
	next *Node
}

type list struct {
	len  int
	head *Node
}

func New(i int) *Node {
	return &Node{
		key:  i,
		next: nil,
	}
}

func (l *list) insertAtBack(i int) {
	newnode := New(i)
	if l.head == nil {
		l.head = newnode
		return
	}
	crr := l.head
	for crr.next != nil {
		crr = crr.next
	}
	crr.next = newnode
	l.len++
	return
}

func (l *list) insertAtFront(i int) {
	newnode := New(i)
	if l.head == nil {
		l.head = newnode
		l.len++
		return
	}

	l.head = newnode
	l.len++
}

func (l *list) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("[%v]-> ", list.key)
		list = list.next
	}
}

func main() {
	l := list{}
	l.insertAtBack(1)
	l.insertAtBack(2)
	l.insertAtBack(3)
	l.Display()
	l.insertAtFront(4)
	l.Display()

}
