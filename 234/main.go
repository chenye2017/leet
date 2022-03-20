package main

import (
	"fmt"
	"math"
)

func main() {
	head := makeList([]int{1, 2})
	fmt.Println(isPalindrome(head))
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

func isPalindrome(head *ListNode) bool {
	count := 1
	now := head
	first := head
	second := head

	for {
		now = now.Next
		if now == nil {
			break
		} else {
			count++
		}
	}

	if count == 1 {
		return true
	}

	before := int(math.Ceil(float64(count) / float64(2)))

	for i := 1; i <= before; i++ {
		first = first.Next
	}

	var prev *ListNode
	for {
		tmp := &ListNode{
			Val:  first.Val,
			Next: prev,
		}
		prev = tmp
		first = first.Next
		if first == nil {
			break
		}
	}

	for i := 1; i <= count/2; i++ {
		firstVal := prev.Val
		secondVal := second.Val
		if firstVal != secondVal {
			fmt.Println(firstVal, secondVal)
			return false
		}
		prev = prev.Next
		second = second.Next
	}
	return true
}
