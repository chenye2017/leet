package main

func main() {

}

type NumArray struct {
	Arr []int
	Sum int
}

func Constructor(nums []int) NumArray {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return NumArray{Arr: nums, Sum: sum}
}

func (this *NumArray) Update(index int, val int) {
	old := this.Arr[index]
	this.Arr[index] = val
	this.Sum = this.Sum - old + val
}

func (this *NumArray) SumRange(left int, right int) int {
	tmp := this.Arr[left : right+1]

	sum := 0
	for _, v := range tmp {
		sum += v
	}
	return sum
}
