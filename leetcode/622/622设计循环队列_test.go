package _22

import "testing"

//执行用时：8 ms, 在所有 Go 提交中击败了96.65%的用户
//内存消耗：6.6 MB, 在所有 Go 提交中击败了96.23%的用户
type MyCircularQueue struct {
	nums       []int
	firstIndex int
	num        int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int, k),
		0,
		0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.num == len(this.nums) {
		return false
	}
	this.nums[(this.firstIndex+this.num)%len(this.nums)] = value
	this.num++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.num > 0 {
		this.firstIndex = (this.firstIndex + 1) % len(this.nums)
		this.num--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.num == 0 {
		return -1
	}
	return this.nums[this.firstIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.num == 0 {
		return -1
	}
	return this.nums[(this.firstIndex+this.num-1)%len(this.nums)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.num == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.num == len(this.nums)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func TestMyCircularQueue(t *testing.T) {
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	obj.DeQueue()
	obj.EnQueue(4)
	obj.Rear()
}
