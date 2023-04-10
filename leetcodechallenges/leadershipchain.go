package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Print() // Output: 1 2 3

}
