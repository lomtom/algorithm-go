package leetcode

import (
	"fmt"
	"testing"
)

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.2 MB,击败了85.00% 的Go用户
func buildArray(target []int, n int) (ans []string) {
	var index int = 0
	for i := 1; i <= n; i++ {
		ans = append(ans, "Push")
		if index >= len(target) || target[index] != i {
			ans = append(ans, "Pop")
		} else {
			index++
		}
	}
	for len(ans) >= 2 {
		if ans[len(ans)-1] == "Pop" && ans[len(ans)-2] == "Push" {
			ans = ans[:len(ans)-2]
		} else {
			break
		}
	}
	return
}

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.2 MB,击败了85.00% 的Go用户
func buildArray1(target []int, n int) (ans []string) {
	var index int = 0
	for i := 1; i <= target[len(target)-1]; i++ {
		if index >= len(target) || target[index] != i {
			ans = append(ans, "Push", "Pop")
		} else {
			ans = append(ans, "Push")
			index++
		}
	}
	return
}

func TestBuildArray(t *testing.T) {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}
