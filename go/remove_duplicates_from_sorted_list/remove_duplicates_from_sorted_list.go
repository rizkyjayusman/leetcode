package main

import (
	"fmt"
)

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}

func main() {
	ln := List{}
	ln.Insert(1)
	ln.Insert(1)
	ln.Insert(2)
	removeDuplicates(ln)
}

func removeDuplicates(ln List) List {
	nln := List{}
	x := -101
	for ln.head != nil {
		if ln.head.info.(int) > x {
			nln.Insert(ln.head.info)
		}
		x = ln.head.info.(int)
		ln.head = ln.head.next
	}
	return nln
}
