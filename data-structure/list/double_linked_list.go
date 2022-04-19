package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

func inserNode(l *Node, data int) *Node {

	list := l

	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	// if lis is null
	if list == nil {
		list = newNode
		return list
	}

	if list.data >= data {
		list.prev = newNode
		newNode.next = list
		return newNode
	}

	for list != nil {
		if list.next != nil && list.next.data >= data {
			fmt.Println("cok")
			newNode := &Node{
				data: data,
				prev: list.prev,
				next: list.next,
			}
			list.next = newNode
			list.next.prev = newNode
			break
		} else if list.next == nil {
			newNode := &Node{
				data: data,
				prev: list,
				next: nil,
			}
			list.next = newNode

			break
		}
		list = list.next
	}

	return l

}

// 1 -> 2 -> 3 =>

func reverse(n *Node) (res *Node) {
	temp := &Node{}
	current := n
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}
	temp.prev.print()
	if temp != nil {
		res = temp.prev
	}

	return res
}

func (n *Node) print() {
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {

	node1 := &Node{
		data: 2,
		prev: nil,
		next: nil,
	}
	node2 := &Node{
		data: 3,
		prev: node1,
		next: nil,
	}
	node1.next = node2
	node3 := &Node{
		data: 4,
		prev: node2,
		next: nil,
	}
	node2.next = node3
	node4 := &Node{
		data: 50,
		prev: node3,
		next: nil,
	}
	node3.next = node4

	h := reverse(node1)
	h.print()
}
