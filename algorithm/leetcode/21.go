package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	resultL := l1
	resultH := l2
	if l1.Val > l2.Val {
		resultL = l2
		resultH = l1
	}
	next := resultL
	for next.Next != nil {
		for resultH != nil {
			if resultH.Val < next.Next.Val {
				temp := next.Next
				next.Next = resultH
				resultH = resultH.Next
				next.Next.Next = temp
				next = next.Next
				continue
			}
			next = next.Next
			break
		}
		if resultH == nil {
			return resultL
		}

	}
	next.Next = resultH
	return resultL
}
