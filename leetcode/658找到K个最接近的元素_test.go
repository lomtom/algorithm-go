package leetcode

import "testing"

func findClosestElements(arr []int, k int, x int) []int {
	sub := func(num1, num2 int) bool {
		var ans1 int
		if num1 > x {
			ans1 = num1 - x
		} else {
			ans1 = x - num1
		}
		var ans2 int
		if num2 > x {
			ans2 = num2 - x
		} else {
			ans2 = x - num2
		}
		return ans1 < ans2 || ans1 == ans2 && num1 < num2
	}
	var ans []int
	for _, num := range arr {
		if k > 0 {
			ans = append(ans, num)
			k--
		} else {
			if sub(ans[0], num) {
				continue
			}
			ans = append(ans[1:], num)
		}
	}
	return ans
}

//执行用时：52 ms, 在所有 Go 提交中击败了11.84%的用户
//内存消耗：6.8 MB, 在所有 Go 提交中击败了7.02%的用户
func TestFindClosestElements(t *testing.T) {
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, 6)
}
