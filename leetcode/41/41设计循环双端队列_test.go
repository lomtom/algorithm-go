package _1

//执行用时：16 ms, 在所有 Go 提交中击败了23.95%的用户
//内存消耗：6.6 MB, 在所有 Go 提交中击败了81.44%的用户
type MyCircularDeque struct {
	nums       []int
	start, end int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		make([]int, k+1),
		0, 0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.start = (this.start - 1 + len(this.nums)) % len(this.nums)
	this.nums[this.start] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.nums[this.end] = value
	this.end = (this.end + 1) % len(this.nums)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % len(this.nums)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.end = (this.end - 1 + len(this.nums)) % len(this.nums)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[this.start]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[(this.end-1+len(this.nums))%len(this.nums)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.start == this.end
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.end+1)%len(this.nums) == this.start
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
