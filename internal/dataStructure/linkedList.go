package dataStructure

import (
"bufio"
"os"
"strconv"
)
var stdout = bufio.NewWriter(os.Stdout)

type node struct{
	next *node
	prev *node
	val int
}
type linkedList struct{
	root *node
	tail *node
}

func (l *linkedList) append(val int){
	if l.root == nil{
		l.root = &node{val:val}
		l.tail = l.root
		return
	}
	l.tail.next = &node{prev:l.tail,val:val}
	l.tail = l.tail.next
}

func (l *linkedList) delete(node *node) {
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

func (l *linkedList) print(){
	defer stdout.Flush()

	node := l.root
	for node.next != nil{
		stdout.WriteString(strconv.Itoa(node.val) +"->")
		node = node.next
	}
	stdout.WriteString(strconv.Itoa(node.val)+"\n")
}

//func main(){
//	list := &linkedList{}
//	list.append(0)
//}
