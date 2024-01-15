package leetcode

import (
	"fmt"
	"testing"
)

// 4
// 1 1 1 1
// 1 2 1
// 1 1 2
// 2 1 1
// 2 2

// 1 : 1
// 2 : 2
// 3 : 3
// 4 : 5
// f(x) = f(x-1) + f(x-2)
func climbStairs(n int) int {
	res := []int{0, 0, 1}
	for i := 0; i < n; i++ {
		res[0] = res[1]
		res[1] = res[2]
		res[2] = res[0] + res[1]
	}
	return res[2]
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了85.78% 的Go用户
func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}
