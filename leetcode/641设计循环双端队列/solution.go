package _1

import (
	"fmt"
	"testing"
)

// 执行用时：16 ms, 在所有 Go 提交中击败了23.95%的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了81.44%的用户
type MyCircularDeque struct {
	nums       []int
	firstIndex int
	num        int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		make([]int, k),
		0, 0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.firstIndex = (this.firstIndex + len(this.nums) - 1) % len(this.nums)
	this.nums[this.firstIndex] = value
	this.num++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	index := (this.firstIndex + this.num) % len(this.nums)
	this.nums[index] = value
	this.num++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.firstIndex = (this.firstIndex + 1) % len(this.nums)
	this.num--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.num--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[this.firstIndex]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[(this.firstIndex+this.num-1)%len(this.nums)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.num == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.num == len(this.nums)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func TestMyCircularDeque(t *testing.T) {
	obj := Constructor(3)
	_ = obj.InsertLast(1)
	_ = obj.InsertLast(2)
	_ = obj.InsertFront(3)
	_ = obj.InsertFront(4)
	fmt.Println(obj.GetFront())
	obj.DeleteLast()
}
