package _656

// 执行耗时:60 ms,击败了100.00% 的Go用户
// 内存消耗:7.2 MB,击败了15.79% 的Go用户
type OrderedStream struct {
	nums  []string
	index int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		make([]string, n),
		0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) (ans []string) {
	this.nums[idKey-1] = value
	for this.index < len(this.nums) {
		if this.nums[this.index] == "" {
			break
		} else {
			ans = append(ans, this.nums[this.index])
		}
		this.index++
	}
	return
}
