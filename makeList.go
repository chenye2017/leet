package main

func main() {
	p := Make([]int,)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Make(arr []int) *ListNode{
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
