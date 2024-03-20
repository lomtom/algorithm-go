package leetcode

const mod int = 1e9 + 7

// 执行耗时:1 ms,击败了33.33% 的Go用户
// 内存消耗:2 MB,击败了16.67% 的Go用户
func minNonZeroProduct(p int) int {
	k := 1<<p - 1
	return k % mod * pow(k-1, p-1) % mod
}

func pow(x, p int) int {
	res := 1
	for x %= mod; p > 0; p-- {
		res = res * x % mod
		x = x * x % mod
	}
	return res
}
