package offer09

type CQueue struct {
	nums []int
}

func Constructor() CQueue {
	return CQueue{
		nums: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.nums = append(this.nums, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.nums) == 0 {
		return -1
	}
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}
