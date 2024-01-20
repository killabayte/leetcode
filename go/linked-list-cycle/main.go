package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {

	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 3}

	result1 := hasCycle(list1)
	fmt.Println("Example 1:", result1)

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 2}
	list2.Next.Next = &ListNode{Val: 3}
	list2.Next.Next.Next = list2.Next

	result2 := hasCycle(list2)
	fmt.Println("Example 2:", result2)
}
