package main

import (
	"errors"
	"fmt"
)

type stack struct {
	values []int
}

var push = func(s *stack, v int) {
	s.values = append(s.values, v)
}
var isEmpty = func(s *stack) bool {
	if len(s.values) == 0 {
		return true
	} else {
		return false
	}
}
var peak = func(s *stack) (int, error) {
	if empty := isEmpty(s); empty {
		return 0, errors.New("is empty")
	} else {
		return s.values[len(s.values)-1], nil
	}
}
var pop = func(s *stack) int {
	tmp := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return tmp
}

type MyQueue struct {
	in  *stack
	out *stack
}

func Constructor() MyQueue {
	m := MyQueue{new(stack), new(stack)}
	return m
}

func (this *MyQueue) Push(x int) {
	push(this.in, x)
}

func (this *MyQueue) Pop() int {
	if !isEmpty(this.out) {
		return pop(this.out)
	}

	for {
		if !isEmpty(this.in) {
			v := pop(this.in)
			push(this.out, v)
		} else {
			break
		}
	}
	return pop(this.out)
}

func (this *MyQueue) Peek() int {
	if !isEmpty(this.out) {
		tmp, _ := peak(this.out)
		return tmp
	}

	for {
		if !isEmpty(this.in) {
			v := pop(this.in)
			push(this.out, v)
		} else {
			break
		}
	}
	tmp, _ := peak(this.out)
	return tmp
}

func (this *MyQueue) Empty() bool {
	return isEmpty(this.out) && isEmpty(this.in)
}

func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Peek())
	fmt.Println(m.Pop())
	fmt.Println(m.Peek())
	fmt.Println(m.Pop())
	fmt.Println(m.Empty())
}
