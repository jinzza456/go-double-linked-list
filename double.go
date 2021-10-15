package main

import (
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

type LinkeList struct {
	root *Node
	tail *Node
}

func (l *LinkeList) AddNode(val int) {
	if l.root == nil { // 맨처음엔 root와 tail 둘다 nil 이기때문에
		l.root = &Node{val: val} // 새로운 노드를 만듦
		l.tail = l.root          // 하나만 있는 상태에서 root와 tail은 같음
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkeList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}
	prev := l.root
	for prev.next != node {
		prev = prev.next
	}
	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
	node.next = nil
}

func (l *LinkeList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkeList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()
	fmt.Printf("tail:%d\n", list.tail.val)

}
