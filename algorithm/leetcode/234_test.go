package main

import (
	"fmt"
	"testing"
)

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow := head
	fast := head

	// find middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// circle
			return false
		}
	}

	reverse := slow
	if fast != nil {
		reverse = slow.Next
	}
	// reverse
	cur := reverse
	for cur.Next != nil {
		temp := reverse
		reverse = cur.Next
		cur.Next = cur.Next.Next
		reverse.Next = temp
	}

	for reverse != nil {
		if head.Val != reverse.Val {
			return false
		}
		reverse = reverse.Next
	}

	return true
}
func TestPalindrome(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println("is palindrome", isPalindrome(head))
}
