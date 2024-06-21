package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.3 MB,击败了100.00% 的Go用户
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	var count int
	var max int
	for i := 0; i < len(temperatureA)-1; i++ {
		if isSame(temperatureA[i+1]-temperatureA[i], temperatureB[i+1]-temperatureB[i]) {
			count++
		} else {
			count = 0
		}
		max = maxFunc(max, count)
	}
	return max
}

func isSame(a, b int) bool {
	if a > 0 && b > 0 {
		return true
	}
	if a < 0 && b < 0 {
		return true
	}
	if a == 0 && b == 0 {
		return true
	}
	return false
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
