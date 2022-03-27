package main

import "fmt"

func main() {
	l := Make([]int{1, 1, 2, 3, 3})
	l = deleteDuplicates(l)
	for {
		if l == nil {
			break
		} else {
			fmt.Println(l.Val)
			l = l.Next
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

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1000,
		Next: head,
	}
	prev := dummy
	for {
		if prev.Next == nil {
			break
		} else {
			if prev.Next.Val == prev.Val {
				prev.Next = prev.Next.Next
			} else {
				prev = prev.Next
			}
		}
	}
	return dummy.Next
}
