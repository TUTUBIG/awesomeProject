package algorithm

/*
				random read		insert/delete
  Array:   			O(1)			O(n)
  Link List:		O(n)			O(1)

  * Array will use a contiguous memory space so that CPU cache could be used
*/

import "fmt"

type ListNode struct {
	next *ListNode
	prev *ListNode
	data interface{}
}

type ListOperation interface {
	print()
	insertN(data interface{}, n int)
	deleteN(n int)
	reverse()
}

type SingleList struct {
	*ListNode
}

type DoubleList struct {
	*ListNode
}

func (sl *SingleList) print() {
	cur := sl.ListNode

	for cur != nil {
		if cur.next == nil {
			fmt.Print(cur.data)
		} else {
			fmt.Print(cur.data, " ->")
		}
		cur = cur.next
	}
}

func (sl *SingleList) insertN(data interface{}, n int) {

}

func (sl *SingleList) reverse() {

}

func (sl *SingleList) checkCircle() bool {
	slow, fast := sl.ListNode, sl.ListNode
	for fast != nil {
		slow = slow.next
		if fast.next == nil {
			return false
		}
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}
	return false
}

func createRoundSingleList(data ...interface{}) SingleList {
	if len(data) == 0 {
		return SingleList{nil}
	}
	listNodes := make([]ListNode, len(data))
	for i := range data {
		listNodes[i].data = data[i]
		if i < len(data)-1 {
			listNodes[i].next = &listNodes[i+1]
		} else {
			listNodes[i].next = &listNodes[0]
		}
	}
	return SingleList{&listNodes[0]}
}

func createSingleList(data ...interface{}) SingleList {
	listNodes := make([]*ListNode, len(data))
	if len(listNodes) < 1 {
		return SingleList{nil}
	}

	for i := 0; i < len(listNodes); i++ {
		listNodes[i] = new(ListNode)
	}

	for i := 0; i < len(listNodes); i++ {
		listNodes[i].data = data[i]
		if i+1 < len(listNodes) {
			listNodes[i].next = listNodes[i+1]
		}
	}

	return SingleList{listNodes[0]}

}

func createDoubleList(data ...interface{}) DoubleList {
	listNodes := make([]*ListNode, len(data))
	if len(listNodes) < 1 {
		return DoubleList{nil}
	}

	for i := 0; i < len(listNodes); i++ {
		listNodes[i] = new(ListNode)
	}

	for i := 0; i < len(listNodes); i++ {
		listNodes[i].data = data[i]
		if i+1 < len(listNodes) {
			listNodes[i].next = listNodes[i+1]
		}
		if i > 0 {
			listNodes[i].prev = listNodes[i-1]
		}
	}

	return DoubleList{listNodes[0]}

}

// todo print
// todo insert
// todo delete
// todo reverse
// todo check circle
// todo find middle node
// todo merge
// todo LRU: Least Recently Used
// todo FIFO: First In, First Out
// todo LFU: Least Frequently Used
