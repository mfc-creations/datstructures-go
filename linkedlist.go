package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	n.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)

		toPrint = toPrint.next
		l.length--
	}
}

func (l *linkedList) deleteNode(num int) {
	if l.length == 0 {
		fmt.Print("Empty list")
	}

	if num == l.head.data {
		fmt.Print("Iam head")
		l.head = l.head.next
		l.length--
		return
	}
	fmt.Print("Not head")
	prev := l.head
	for prev.next.data != num {
		if prev.next.next == nil {
			fmt.Print("Value not found")
			return
		}
		prev = prev.next
	}

	prev.next = prev.next.next
	l.length--

}

// func main() {
// 	myList := linkedList{}
// 	node1 := &node{data: 8}
// 	node2 := &node{data: 10}
// 	node3 := &node{data: 23}

// 	myList.prepend(node1)
// 	myList.prepend(node2)
// 	myList.prepend(node3)

// 	myList.printListData()
// 	myList.deleteNode(223)
// 	myList.printListData()

// }
