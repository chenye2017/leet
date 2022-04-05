package main

import "fmt"

func main() {
	l := Constructor(3)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)                 // 4, 3 ,2
	fmt.Println(l.Get(4), "sh") // 4, 3 ,2
	fmt.Println(l.Get(3), "sh") // 3, 4, 2
	fmt.Println(l.Get(2), "sh") // 2, 3, 4
	fmt.Println(l.Get(1), "sh")
	l.Put(5, 5) // 5, 2, 3
	//fmt.Println(l.M, "----")
	//panic("----")
	fmt.Println(l.Get(1), "sh")
	fmt.Println(l.Get(2), "sh")
	fmt.Println(l.Get(3), "sh")
	fmt.Println(l.Get(4), "sh")
	fmt.Println(l.Get(5), "sh")
}

// 双链表
// 头尾指针
// key
// val
// address
// 双指针的写法
// 删除节点：移动2个指针
// 增加节点：移动4个指针
type LRUCache struct {
	Cap  int
	M    map[int]*LRUCacheNode
	Head *LRUCacheNode
	Tail *LRUCacheNode
}

type LRUCacheNode struct {
	Key  int
	Val  int
	Next *LRUCacheNode
	Prev *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	headDummy := &LRUCacheNode{
		Key:  0,
		Val:  0,
		Next: nil,
		Prev: nil,
	}
	endDummy := &LRUCacheNode{
		Key:  0,
		Val:  0,
		Next: nil,
		Prev: nil,
	}

	headDummy.Next = endDummy
	endDummy.Prev = headDummy

	return LRUCache{
		Cap:  capacity,
		M:    make(map[int]*LRUCacheNode),
		Head: headDummy,
		Tail: endDummy,
	}
}

func (this *LRUCache) Get(key int) int {
	//fmt.Println(key, this.Head.Next.Val, this.Head.Next.Next.Val, this.Head.Next.Next.Next.Val, "--get-")
	//fmt.Println(key, this.Tail.Prev.Val, this.Tail.Prev.Prev.Val, this.Tail.Prev.Prev.Prev.Val, "--back-")
	if tmp, ok := this.M[key]; ok {
		if len(this.M) > 1 {
			//	fmt.Println(this.Tail.Val, this.Tail.Prev.Val, tmp.Next.Val, "--1")
			// 删除元素
			tmp.Prev.Next = tmp.Next
			tmp.Next.Prev = tmp.Prev
			//	fmt.Println(tmp.Next.Val, this.Tail.Prev.Val, "--2")

			tmp.Next = this.Head.Next
			tmp.Prev = this.Head
			this.Head.Next.Prev = tmp
			this.Head.Next = tmp

			//	fmt.Println(key, this.Tail.Prev.Val, this.Tail.Prev.Prev.Val, this.Tail.Prev.Prev.Prev.Val, "--backde-")
		}
		return tmp.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 判断元素是否存在，存在要删除
	var tmp *LRUCacheNode
	tmp, _ = this.M[key]

	//
	var isDelete bool
	if tmp != nil {
		// 是旧的元素，就不用删除
		isDelete = false
	} else {
		// 不是旧的，就可能要删除
		if len(this.M) < this.Cap {
			isDelete = false
		} else {
			isDelete = true
		}
	}

	// 直接添加, 头元素
	Node := &LRUCacheNode{
		Key:  key,
		Val:  value,
		Next: this.Head.Next,
		Prev: this.Head,
	}

	this.M[key] = Node
	this.Head.Next.Prev = Node
	this.Head.Next = Node

	if this.Tail.Prev.Prev != nil {
		//	fmt.Println(isDelete, this.M, this.Head.Next.Val, this.Tail.Prev.Val, this.Tail.Prev.Prev.Val, "--put-start-", value)
	}

	if this.Tail.Prev == this.Head {
		this.Tail.Prev = Node
	}

	this.M[key] = Node

	//
	if tmp != nil {
		// 删除元素
		tmp.Prev.Next = tmp.Next
		tmp.Next.Prev = tmp.Prev
	}
	/*if this.M[1] != nil {
		fmt.Println(this.M[1].Prev.Val, "--1--")
	}
	if this.M[2] != nil {
		fmt.Println(this.M[2].Prev.Val, "--2--")
	}*/

	// 删除尾巴元素
	if isDelete {
		delete(this.M, this.Tail.Prev.Key)
		this.Tail.Prev.Prev.Next = this.Tail
		this.Tail.Prev = this.Tail.Prev.Prev
	}
	//fmt.Println(isDelete, this.M, this.Head.Next.Val, this.Tail.Prev.Val, "--put-", value)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
