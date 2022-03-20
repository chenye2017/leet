package main

import "fmt"

func main() {
	head := makeList([]int{1, 2})
	r := removeNthFromEnd(head, 1)
	r.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(arr []int) *ListNode {

	var head *ListNode
	now := head
	for _, v := range arr {
		tmp := &ListNode{
			Val:  v,
			Next: nil,
		}
		if now == nil {
			now = tmp
			head = tmp
		} else {
			now.Next = tmp
			now = now.Next
		}
	}

	return head
}

func (l *ListNode) String() {
	for {
		if l != nil {
			fmt.Println(l.Val)
			l = l.Next
		} else {
			break
		}
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dummy

	first := head
	second := head
	for i := 1; i < n; i++ {
		first = first.Next
	}

	for {
		first = first.Next

		if first == nil {
			break
		}
		prev = second
		second = second.Next
	}
	prev.Next = second.Next

	return dummy.Next
}
