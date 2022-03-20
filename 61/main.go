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
	head := makeList([]int{1})
	res3 := rotateRight(head, 1)
	res3.String()
	fmt.Println("===============")
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}

	// 遍历 拿到 链表len
	var count int
	var end *ListNode
	var res *ListNode
	m := make(map[int]*ListNode)
	lenIndex := head
	for {
		if lenIndex != nil {
			count++
			m[count] = lenIndex
			end = lenIndex
			lenIndex = lenIndex.Next
		} else {
			break
		}
	}
	// 获取k
	if k >= count {
		k = k % count
	}
	if k == 0 {
		return head
	}

	// 第k个元素末尾指向nil
	res = m[count-k].Next
	m[count-k].Next = nil
	end.Next = head
	return res
}
