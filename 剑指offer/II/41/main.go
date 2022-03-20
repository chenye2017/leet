package main

import "fmt"

type MovingAverage struct {
	values []int
	size   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	m := MovingAverage{size: size}
	return m
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.values) >= this.size {
		this.values = this.values[1:]
		this.values = append(this.values, val)
	} else {
		this.values = append(this.values, val)
	}
	sum := 0
	for _, v := range this.values {
		sum += v
	}
	return float64(sum) / float64(len(this.values))
}

func main() {
	m := Constructor(3)
	fmt.Println(m.Next(1))
	fmt.Println(m.Next(10))
	fmt.Println(m.Next(3))
	fmt.Println(m.Next(5))
}
