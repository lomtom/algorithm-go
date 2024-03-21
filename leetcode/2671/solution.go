package leetcode

// 执行耗时:320 ms,击败了100.00% 的Go用户
// 内存消耗:65.1 MB,击败了60.00% 的Go用户

type FrequencyTracker struct {
	countMap  map[int]int
	numberMap map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		countMap:  make(map[int]int),
		numberMap: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	this.countMap[this.numberMap[number]]--
	this.numberMap[number]++
	this.countMap[this.numberMap[number]]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.numberMap[number] == 0 {
		return
	}
	this.countMap[this.numberMap[number]]--
	this.numberMap[number]--
	this.countMap[this.numberMap[number]]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.countMap[frequency] > 0
}
