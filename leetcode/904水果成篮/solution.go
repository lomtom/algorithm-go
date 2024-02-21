package leetcode

// 执行耗时:100 ms,击败了64.60% 的Go用户
// 内存消耗:8.3 MB,击败了72.20% 的Go用户
func totalFruit(fruits []int) (ans int) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	cnt := map[int]int{}
	left := 0
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
