package main

import "fmt"

func main() {
	l := Make([]int{1})
	mid := middleNode(l)
	fmt.Println(mid.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Make(arr []int) *ListNode {
	count := len(arr)
	var prev *ListNode
	for i := count - 1; i >= 0; i-- {
		prev = &ListNode{
			Val:  arr[i],
			Next: prev,
		}
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	if fast.Next == nil {
		return slow
	}

	for {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		} else {
			fast = fast.Next
			if fast == nil || fast.Next == nil {
				break
			}
		}
	}
	return slow
}
