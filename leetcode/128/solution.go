package leetcode

// 执行耗时:142 ms,击败了22.20% 的Go用户
// 内存消耗:13.7 MB,击败了5.02% 的Go用户
func longestConsecutive(nums []int) int {
	var numSet = make(map[int]int)
	for _, num := range nums {
		numSet[num] = num
	}
	var find func(num int) (int, bool)
	find = func(num int) (int, bool) {
		if v, ok := numSet[num]; ok {
			if v != num {
				if v1, ok1 := find(v); ok1 {
					numSet[num] = v1
				}
			}
			return numSet[num], true
		}
		return 0, false
	}
	var union func(num1, num2 int)
	union = func(num1, num2 int) {
		v1, ok1 := find(num1)
		v2, ok2 := find(num2)
		if ok1 && ok2 {
			if v1 != v2 {
				numSet[v1] = v2
			}
		}
	}
	for _, num := range nums {
		union(num, num+1)
	}
	var maxLen = 0
	for _, num := range nums {
		if v, ok := find(num); ok {
			if v-num+1 > maxLen {
				maxLen = v - num + 1
			}
		}
	}
	return maxLen
}
