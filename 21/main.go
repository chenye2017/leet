package main

import "fmt"

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

func main() {
	head1 := makeList([]int{1, 2, 4})
	head2 := makeList([]int{1, 3, 4})

	res := mergeTwoLists(head1, head2)
	res.String()
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode
	var prev *ListNode
	list1Head := list1
	list2Head := list2

	for {
		now := new(ListNode)
		val1 := list1Head.Val
		val2 := list2Head.Val
		if val1 > val2 {
			now.Val = val2
			list2Head = list2Head.Next
		} else {
			now.Val = val1
			list1Head = list1Head.Next
		}

		if prev != nil {
			prev.Next = now
			prev = now
		} else {
			res = now
			prev = now
		}

		if list1Head == nil {
			now.Next = list2Head
			break
		}
		if list2Head == nil {
			now.Next = list1Head
			break
		}
	}
	return res
}
