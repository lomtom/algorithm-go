package leetcode

// 执行耗时:3 ms,击败了17.65% 的Go用户
// 内存消耗:2.2 MB,击败了100.00% 的Go用户
func minimumOperations(num string) int {
	var zeroCount, fiveCount, otherCount int
	for i := len(num) - 1; i >= 0; i-- {
		switch num[i] {
		case '0':
			if zeroCount > 0 {
				return zeroCount - 1 + fiveCount + otherCount
			}
			zeroCount++
		case '5':
			if zeroCount > 0 {
				return zeroCount - 1 + fiveCount + otherCount
			}
			fiveCount++
		case '2', '7':
			if fiveCount > 0 {
				return zeroCount + fiveCount - 1 + otherCount
			}
			otherCount++
		default:
			otherCount++
		}
	}
	if zeroCount > 0 {
		return zeroCount - 1 + fiveCount + otherCount
	}
	return zeroCount + fiveCount + otherCount
}
