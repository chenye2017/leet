package main

import "fmt"

type MyStack struct {
	tmp []int
	use []int
}

func Constructor() MyStack {
	s := MyStack{
		tmp: nil,
		use: nil,
	}
	return s
}

func (this *MyStack) Push(x int) {
	if len(this.use) > 0 {
		for _, v := range this.use {
			this.tmp = append(this.tmp, v)
		}
		this.use = []int{x}
		for _, v := range this.tmp {
			this.use = append(this.use, v)
		}
		this.tmp = []int{}
	} else {
		this.use = append(this.use, x)
	}
}

func (this *MyStack) Pop() int {
	tmp := this.use[0]
	this.use = this.use[1:]
	return tmp
}

func (this *MyStack) Top() int {
	return this.use[len(this.use)-1]
}

func (this *MyStack) Empty() bool {
	if len(this.use) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Top())
	fmt.Println(m.Pop())
	fmt.Println(m.Empty())
}
