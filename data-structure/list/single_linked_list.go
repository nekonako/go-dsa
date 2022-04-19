package main

import (
	"fmt"
	"log"
)

// single linked list
type Node struct {
	data int
	next *Node
}

// add node at the end of linked list
func (n *Node) append(data int) {

	new_node := &Node{
		data: data,
		next: nil,
	}

	if n == nil {
		n = new_node
		return
	}

	last := n

	// looping to last node
	for last.next != nil {
		last = last.next
	}

	// move next node of last node to new node
	last.next = new_node

}

func (n *Node) print() {
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

// add node at first of linked list
func (n *Node) push(data int) {

	// take value of receiverRepresentation
	var head Node = *n

	new_node := Node{
		data: data,
		next: &head,
	}

	// dereference receiver to new_node pointer
	*n = new_node

}

func (n *Node) insertAfter(prev_node *Node, data int) {

	if prev_node == nil {
		log.Fatal("prev node cannot be null")
		return
	}

	new_node := Node{
		data: data,
		next: prev_node.next,
	}

	prev_node.next = &new_node

}

func (n *Node) delete(key int) {

	var next_node Node

	for n != nil && n.data != key {

		n = n.next
		if n.next == nil {
			continue
		}

		// if n.data == key, skip this node and move next node to next_node variable
		next_node = *n.next

	}

	*n = next_node

}

func (n Node) length() (result int) {
	for &n != nil {
		result += 1
		n = *n.next
	}
	return
}

func (n *Node) recursiveLength() (result int) {
	if n != nil {
		result += 1
		n = n.next
	}
	return
}

func (n *Node) insertPostition(data int, postion int) {

	new_node := &Node{
		data: data,
		next: nil,
	}

	l := n
	counter := 0
	var prev *Node

	for l != nil {
		counter += 1
		if counter == postion {
			new_node.next = l
			prev.next = new_node
			break
		}

		prev = l
		l = l.next
	}
	n.print()
	//	prev.print()
}

func (n Node) reverse() (node Node) {

	current := &n
	var prev *Node = nil
	var next *Node = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	return *prev

}

func main() {

	var list Node

	list.append(1)
	list.append(2)
	list.append(3)

	//list.insertPostition(4, 3)
	l := list.recursiveLength()
	fmt.Println(l)
	//reversed := list.reverse()
	//reversed.print()

}
