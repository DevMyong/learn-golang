package dataStructure

import (
"bufio"
"os"
"strconv"
)
var stdout = bufio.NewWriter(os.Stdout)

type Node struct{
	next *Node
	prev *Node
	val int
}
type LinkedList struct{
	root *Node
	tail *Node
}

func (l *LinkedList) Append(val int){
	if l.root == nil{
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{prev: l.tail,val:val}
	l.tail = l.tail.next
}

func (l *LinkedList) Delete(node *Node) {
	if node == l.root{
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}
	prev := node.prev

	if node == l.tail{
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	}else {
		prev.next = prev.next.next
		node.next.prev = prev
		node.prev, node.next = nil, nil
	}
}

func (l *LinkedList) Print(){
	defer stdout.Flush()

	node := l.root
	for node.next != nil{
		stdout.WriteString(strconv.Itoa(node.val) +"->")
		node = node.next
	}
	stdout.WriteString(strconv.Itoa(node.val)+"\n")
}

//func main(){
//	list := &LinkedList{}
//	list.Append(0)
//}
