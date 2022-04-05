package main

import (
	"fmt"
)

func main() {
	s, end, m := Make([]int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5})
	end.Next = m[6]

	c := detectCycle(s)
	fmt.Println(c.Val, "--")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Make(arr []int) (*ListNode, *ListNode, map[int]*ListNode) {
	count := len(arr)
	var prev *ListNode

	var end *ListNode

	m := make(map[int]*ListNode)
	for i := count - 1; i >= 0; i-- {

		prev = &ListNode{
			Val:  arr[i],
			Next: prev,
		}

		if i == count-1 {
			end = prev
		}
		m[i] = prev
	}
	return prev, end, m
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow, slowSlow := head, head, head

	if slow == nil {
		return nil
	}

	for {
		//	fmt.Println(slow.Val, fast.Val, "-----")
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}

		slow = slow.Next

		// 相遇处
		if slow == fast {
			//	fmt.Println(slow.Val, "----xiangyu")
			if slow == slowSlow {
				return slow
			}

			for {
				fast = fast.Next
				slowSlow = slowSlow.Next
				//		fmt.Println(fast.Val, slowSlow.Val, "--loop")

				if slowSlow == fast {
					return slowSlow
				}
			}
		}
	}

}
