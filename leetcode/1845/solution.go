package leetcode

// 执行耗时:853 ms,击败了8.70% 的Go用户
// 内存消耗:30.4 MB,击败了30.44% 的Go用户

type SeatManager struct {
	stack []bool
	first int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		make([]bool, n),
		0,
	}
}

func (this *SeatManager) Reserve() int {
	this.stack[this.first] = true
	var num = this.first + 1
	for i := this.first; i < len(this.stack); i++ {
		this.first = i
		if !this.stack[i] {
			break
		}
	}
	return num
}

func (this *SeatManager) Unreserve(seatNumber int) {
	index := seatNumber - 1
	this.stack[index] = false
	if index < this.first {
		this.first = index
	}
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
