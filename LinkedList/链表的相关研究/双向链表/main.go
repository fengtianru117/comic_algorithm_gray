package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// NewHeadNode 生成头节点
func NewHeadNode() *Node {
	return &Node{0, nil}
}

// Traverse 遍历链表
func (n *Node) Traverse() {
	nextNode := n.Next
	for nextNode != nil {
		fmt.Println(nextNode.Data)
		nextNode = nextNode.Next
	}
	fmt.Println("----------------Done!----------------")
}

// NewNext 生成新的子节点
func (n *Node) NewNext(data int) *Node {
	next := &Node{data, nil}
	n.Next = next
	return next
}

// //  Generate 通过数组生成链表
// func (n *Node) Generate(s []int) {
// 	for _, v := range s {
// 		nextNode := n.NewNext(v)
// 		nextNode = nextNode.Next
// 	}
// }

// InitNode 初始化节点
func (n *Node) InitNode(data int) {
	n.Data = data
}

// Generate 通过数组生成链表
func Generate(head *Node, s []int) {
	var next *Node
	next = head.NewNext(s[0])
	for _, v := range s[1:] {
		next = next.NewNext(v)
	}
}

func main() {
	// s := []int{1, 2, 3, 4, 5}
	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}
	head := NewHeadNode()
	// head.Generate(s)
	// next := head.NewNext(1)
	// next = next.NewNext(2)
	// next = next.NewNext(3)
	// fmt.Println(next3, "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	Generate(head, s)
	head.Traverse()
}
