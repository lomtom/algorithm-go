package leetcode

// 执行耗时:24 ms,击败了100.00% 的Go用户
// 内存消耗:6.7 MB,击败了61.54% 的Go用户
func maximumBinaryString(binary string) string {
	var zeroCount int
	var firstZeroIndex int
	for index := len(binary) - 1; index >= 0; index-- {
		if binary[index] == '0' {
			zeroCount++
			firstZeroIndex = index
		}
	}
	if zeroCount == 0 {
		return binary
	}
	result := make([]byte, len(binary))
	for index := 0; index < len(binary); index++ {
		if index == zeroCount+firstZeroIndex-1 {
			result[index] = '0'
		} else {
			result[index] = '1'
		}
	}
	return string(result)
}
