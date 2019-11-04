package main

import "fmt"

// Node is list node
type Node struct {
	Data int
	Next *Node
}

// List is single list
type List struct {
	HeadNode *Node
}

// NewNode create a new node
func NewNode(value int) *Node {
	return &Node{value, nil}
}

// NewList
func NewList() *List {
	return &List{&Node{0, nil}}
}

// Append insert a new node after the list
func (list *List) Append(node *Node) {
	tmpList := list.HeadNode
	for {
		if tmpList.Next == nil {
			break
		} else {
			tmpList = tmpList.Next
		}
	}
	tmpList.Next = node
}

// PrintList print the details of the list
func (list *List) PrintList() {
	tmpList := list.HeadNode.Next
	i := 0
	for ; ; i++ {
		if tmpList.Next == nil {
			break
		} else {
			fmt.Printf("Node: %d , Value: %d\n", i+1, tmpList.Data)
			tmpList = tmpList.Next
		}
	}
	fmt.Println("#####################")
	fmt.Printf("Node: %d , Value: %d\n", i+1, tmpList.Data)
	fmt.Println("#####################")
}

func main() {

	var list *List
	list = NewList()

	for i := 0; i < 1000; i++ {
		node := NewNode(i)
		list.Append(node)
	}
	list.PrintList()

}
