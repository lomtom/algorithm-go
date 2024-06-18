package leetcode

var m = map[int]map[int]bool{
	1: {
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	},
	2: {
		1: true,
		3: true,
		5: true,
		7: true,
		9: true,
	},
	3: {
		1: true,
		2: true,
		4: true,
		5: true,
		7: true,
		8: true,
	},
	4: {
		1: true,
		3: true,
		5: true,
		7: true,
		9: true,
	},
	5: {
		1: true,
		2: true,
		3: true,
		4: true,
		6: true,
		7: true,
		8: true,
		9: true,
	},
	6: {
		1: true,
		5: true,
		7: true,
	},
	7: {
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		8: true,
		9: true,
	},
	8: {
		1: true,
		3: true,
		5: true,
		7: true,
		9: true,
	},
	9: {
		1: true,
		2: true,
		4: true,
		5: true,
		7: true,
		8: true,
	},
}

// 执行耗时:30 ms,击败了8.70% 的Go用户
// 内存消耗:4.2 MB,击败了100.00% 的Go用户
func countBeautifulPairs(nums []int) (result int) {
	for i := 0; i < len(nums); i++ {
		var tempI = nums[i]
		for tempI >= 10 {
			tempI = tempI / 10
		}
		for j := i + 1; j < len(nums); j++ {
			tempJ := nums[j] % 10
			if _, ok := m[tempI][tempJ]; ok {
				result++
			}
		}
	}
	return
}
