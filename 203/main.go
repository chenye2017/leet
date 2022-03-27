package main

import "fmt"

func main() {
	p := Make([]int{})
	h := removeElements(p, 1)

	for {
		if h == nil {
			break
		} else {
			fmt.Println(h.Val)
			h = h.Next
		}
	}
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	index := dummy.Next
	prev := dummy

	for {
		if index == nil {
			break
		}
		if index.Val == val {
			prev.Next = index.Next // prev 不移动
		} else {
			prev = prev.Next
		}
		index = index.Next
	}

	return dummy.Next
}
