package leetcode

func minimumPossibleSum1(n int, target int) int {
	const mod = 1e9 + 7
	m := make(map[int]bool)
	count := 0
	x := 1
	for count < n {
		if !m[target-x] {
			m[x] = true
			count++
		}
		x++
	}
	var sum int
	for key := range m {
		sum += key
	}
	return sum % mod
}

// (m*(m+1) + (target+n-m-1)*(target+n-m) - (target-1)*(target))/2
// (m*m + m + target * target + target*n - target*m + target*n + n*n - n*m - target*m - n*m + m*m - target - n + m - target * target + target)/2
// (2*m*m + 2*m + 2*target*n - 2*target*m + n*n - 2*n*m - n)/2
// m*m + m + target*n - target*m - n*m + (n*n-n)/2
// 执行耗时:0 ms,击败了100.00% 的Go用
// 内存消耗:2 MB,击败了18.18% 的Go用户
func minimumPossibleSum(n int, target int) int {
	const mod int = 1e9 + 7
	m := target / 2
	if n <= m {
		return ((1 + n) * n / 2) % mod
	}
	return (m*m + m + target*n - target*m - n*m + (n*n-n)/2) % mod
}
