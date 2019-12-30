package algorithm

import "fmt"

type ListNode struct {
	next *ListNode
	prev *ListNode
	data interface{}
}

type SingleList struct {
	*ListNode
}

type DoubleList struct {
	*ListNode
}

func (sl *SingleList) Println() {
	if sl == nil {
		fmt.Println("nil")
		return
	}

	if sl.CheckCircle() {
		fmt.Println("circle list")
		return
	}

	cur := sl.ListNode

	for cur != nil {
		fmt.Print(cur.data)
		if cur.next != nil {
			fmt.Print("->")
		}
		cur = cur.next
	}
	fmt.Print("\n")
}

func (sl *SingleList) Reverse() {
	if sl == nil {
		return
	}
	if sl.CheckCircle() {
		fmt.Println("circle list")
		return
	}
	cur := sl.ListNode
	var curOld *ListNode
	for cur != nil {
		temp := cur.next
		cur.next = curOld
		curOld = cur
		cur = temp
	}
	sl.ListNode = curOld
}

func (sl *SingleList) GenerateCircle() {
	if sl == nil {
		return
	}

	head, cur := sl.ListNode, sl.ListNode

	for cur != nil {
		if cur.next == nil {
			cur.next = head
			break
		}
		cur = cur.next
	}
}

func (sl *SingleList) CheckCircle() bool {
	if sl == nil || sl.next == nil {
		return false
	}

	cur, curFast := sl.ListNode, sl.next

	for curFast != nil {
		if cur == curFast {
			return true
		}
		cur = cur.next
		if curFast.next == nil {
			return false
		}
		curFast = curFast.next.next
	}
	return false
}

func (dl *DoubleList) Println() {
	if dl == nil {
		fmt.Println("nil")
	}

	cur := dl.ListNode

	for cur != nil {
		fmt.Print(cur.data)
		if cur.next != nil {
			fmt.Print("->")
		}
		cur = cur.next
	}
	fmt.Print("\n")
}

func (dl *DoubleList) Reverse() {
	if dl == nil {
		return
	}
	cur := dl.ListNode

	for cur != nil {
		temp := cur.prev
		cur.prev = cur.next
		cur.next = temp
		dl.ListNode = cur
		cur = cur.prev
	}
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
