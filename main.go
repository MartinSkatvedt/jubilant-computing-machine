package main

import "fmt"

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func createNode(value int) *Node {
	return &Node{value: value, next: nil, previous: nil}
}

func (head *Node) insertEnd(node *Node) {
	next := head

	for next.next != nil {
		next = next.next
	}

	next.next = node
	node.previous = next
}

func (head *Node) insertStart(node *Node) *Node {
	head.previous = node
	node.next = head
	return node
}

func (indexNode *Node) insertAfter(node *Node) {
	node.previous = indexNode
	node.next = indexNode.next
	indexNode.next.previous = node
	indexNode.next = node
}

func (head *Node) deleteNode(node *Node) *Node {

	if node == head {
		node.next.previous = nil
		return node.next
	} else if node.next == nil {
		node.previous.next = nil
		return head
	} else {
		node.previous.next = node.next
		node.next.previous = node.previous
		return head
	}
}

func (head *Node) displayList() {

	fmt.Println("Head = ", head.value)
	next := head.next
	counter := 1
	for next != nil {
		fmt.Println("Node", counter, "=", next.value)
		counter++
		next = next.next
	}
}

func main() {
	head := createNode(1)

	newNode := createNode(5)

	head.insertEnd(newNode)

	head.displayList()
	fmt.Println("")

	head = head.deleteNode(newNode)

	head.displayList()
}
