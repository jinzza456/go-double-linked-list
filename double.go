package main

import (
	"fmt"
)

type Node struct {
	next *Node // 다음 노드를 기억한다.
	prev *Node // 전의 노드를 기억한다.
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
	l.tail.next = &Node{val: val} // tail 의 다음 노드를 만듦
	prev := l.tail                // 이전 tail을 기억함
	l.tail = l.tail.next          // tail을 다음노드로 넣어서 제일 끝의 노드를 항상 tail로 유지
	l.tail.prev = prev            // 새로운 tail의 전노드를 prev로 기억함
}

func (l *LinkeList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}
	prev := node.prev // 지우고자하는 로그의 이전으로

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
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

func (l *LinkeList) PrintReverse() {
	node := l.tail         // tail을 시작점으로
	for node.prev != nil { //tail의 이전이 없을때까지 진행
		fmt.Printf("%d ->", node.val) //val값 출력
		node = node.prev              //node를 이전 노드로 초기화
	}
	fmt.Printf("%d\n", node.val) // 현제 노드 출력
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

	list.PrintReverse()
	fmt.Printf("tail:%d\n", list.tail.val)

}
