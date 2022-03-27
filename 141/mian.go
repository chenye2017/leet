package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
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

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	quickNow := dummy
	slowNow := dummy

	for {
		quickNow = quickNow.Next
		if quickNow == nil {
			return false
		}
		quickNow = quickNow.Next
		if quickNow == nil {
			return false
		}

		slowNow = slowNow.Next
		if quickNow == slowNow {
			return true
		}
	}
}
