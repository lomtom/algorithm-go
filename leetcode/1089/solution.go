package leetcode

// 执行耗时:4 ms,击败了50.00% 的Go用户
// 内存消耗:3.4 MB,击败了100.00% 的Go用户
func duplicateZeros(arr []int) {
	var left, right = -1, -1
	for index := range arr {
		if arr[index] == 0 {
			right++
		}
		left++
		right++
	}
	for left >= 0 {
		if right < len(arr) {
			arr[right] = arr[left]
		}
		right--
		if arr[left] == 0 {
			if right < len(arr) {
				arr[right] = arr[left]
			}
			right--
		}
		left--
	}
}
