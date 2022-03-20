package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var head *ListNode
	now := head
	for _, v := range arr {
		tmp := &ListNode{
			Val:  v,
			Next: nil,
		}
		if now == nil {
			now = tmp
			head = now
		} else {
			now.Next = tmp
			now = now.Next
		}

	}
	head.String()
	head = reverseList(head)
	head.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var res *ListNode
	var prev *ListNode
	now := head
	for {
		tmp := &ListNode{
			Val:  now.Val,
			Next: nil,
		}
		tmp.Next = prev
		prev = tmp

		if now.Next == nil {
			res = prev
			break
		} else {
			now = now.Next
		}
	}
	return res
}
