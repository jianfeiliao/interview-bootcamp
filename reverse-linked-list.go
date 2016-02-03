package main

import "fmt"

func main() {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	node4 := &node{value: 4}
	node5 := &node{value: 5}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	printList(node1)
	reverse(node1)
	printList(node5)
}

type node struct {
	value int
	next  *node
}

func printList(n *node) {
	for n != nil {
		fmt.Print(n.value, "=>")
		n = n.next
	}
	fmt.Println("nil")
}

func reverse(n *node) {
	if n.next == nil {
		return
	}

	curr := n
	next := n.next

	reverse(next)
	//fmt.Println("set ", n2.value, " next => ", n1.value)
	next.next = curr
	curr.next = nil
}

func reverse2(n *node) *node {
	var curr, prev, next *node
	curr = n

	for curr != nil {
		// points to the next node
		next = curr.next

		// reverse the current pointer to previous node
		curr.next = prev

		// step forward
		prev = curr
		curr = next
	}
	return prev
}
