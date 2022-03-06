package main

import (
	"fmt"
)

func main() {
	/*l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l3,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l4,
	}

	l4_a := &ListNode{
		Val:  4,
		Next: nil,
	}
	l6_a := &ListNode{
		Val:  6,
		Next: l4_a,
	}
	l5_a := &ListNode{
		Val:  5,
		Next: l6_a,
	}*/

	l9_1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l9_2 := &ListNode{
		Val:  9,
		Next: l9_1,
	}

	l9_3 := &ListNode{
		Val:  9,
		Next: l9_2,
	}
	l9_4 := &ListNode{
		Val:  9,
		Next: l9_3,
	}

	l9_5 := &ListNode{
		Val:  9,
		Next: l9_4,
	}
	l9_6 := &ListNode{
		Val:  9,
		Next: l9_5,
	}

	l9_7 := &ListNode{
		Val:  9,
		Next: l9_6,
	}

	l9_4_1 := &ListNode{
		Val:  9,
		Next: nil,
	}

	l9_3_1 := &ListNode{
		Val:  9,
		Next: l9_4_1,
	}
	l9_2_1 := &ListNode{
		Val:  9,
		Next: l9_3_1,
	}

	l9_1_1 := &ListNode{
		Val:  9,
		Next: l9_2_1,
	}

	res := addTwoNumbers(l9_7, l9_1_1)
	res.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() {
	for {
		if l == nil {
			return
		} else {
			fmt.Println(l.Val)
			l = l.Next
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	var index *ListNode

	jin := 0
	for {
		tmp1 := 0
		tmp2 := 0
		if l1 != nil {
			tmp1 = l1.Val
		}
		if l2 != nil {
			tmp2 = l2.Val
		}

		v := tmp1 + tmp2

		if res == nil {
			res = &ListNode{
				Val:  (v + jin) % 10,
				Next: nil,
			}
			jin = (v + jin) / 10
			index = res
		} else {
			index.Next = &ListNode{
				Val:  (v + jin) % 10,
				Next: nil,
			}
			jin = (v + jin) / 10
			index = index.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if jin == 0 {
				return res
			} else {
				index.Next = &ListNode{
					Val:  jin,
					Next: nil,
				}
				return res
			}
		}
	}

}
