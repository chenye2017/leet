package main

import (
	"fmt"
)

func main() {
	head := makeList([]int{5})
	res := reverseBetween(head, 1, 1)
	res.String()
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	prev := &ListNode{
		Val:  0,
		Next: head,
	}

	dummy := prev
	index := 1
	now := head

	var tmp *ListNode // 翻转开始
	var start *ListNode
	for {
		if index == left {
			tmp = &ListNode{
				Val:  now.Val,
				Next: nil,
			}
			start = tmp
		}
		if index < left {
			prev = prev.Next
		}
		if index > left {
			tmp = &ListNode{
				Val:  now.Val,
				Next: tmp,
			}
		}
		index++
		// 指针还没移动
		if index > right {
			// prev
			start.Next = now.Next
			prev.Next = tmp
			break
		}
		now = now.Next
	}
	return dummy.Next
}
