package _33

// 执行耗时:92 ms,击败了99.27% 的Go用户
// 内存消耗:8.1 MB,击败了51.82% 的Go用户
type RecentCounter struct {
	nums []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		[]int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.nums = append(this.nums, t)
	temp := t - 3000
	for len(this.nums) > 0 && this.nums[0] < temp {
		this.nums = this.nums[1:]
	}
	return len(this.nums)
}
