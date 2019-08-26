package main

import "fmt"

type SingleListNode struct {
	next *SingleListNode
	data interface{}
}

type SingleList struct {
	head *SingleListNode
}

func (sl *SingleList) Print() {
	if sl.head == nil {
		fmt.Println("nil")
	}
	cur := sl.head
	for cur != nil {
		fmt.Printf("->%+v",cur.data)
		cur = cur.next
	}
}

func createSingleList(n int) SingleList {
	var singleList SingleList
	singleListNodes := make([]SingleListNode,n)
	for i:=0;i<len(singleListNodes)-1;i++{
		singleListNodes[i].data = i
		singleListNodes[i].next = &singleListNodes[i+1]
	}
	singleListNodes[len(singleListNodes)-1].data = len(singleListNodes)-1
	singleListNodes[len(singleListNodes)-1].next = nil
	singleList.head = &singleListNodes[0]
	return singleList
}

func main()  {
	singleList := createSingleList(5)
	singleList.Print()
}
