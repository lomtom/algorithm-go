package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minimumAddedInteger([]int{3, 0, 3, 5, 3}, []int{10, 10, 10}))                        // 7
	t.Log(minimumAddedInteger([]int{4, 6, 3, 1, 4, 2, 10, 9, 5}, []int{5, 10, 3, 2, 6, 1, 9})) // 0
	t.Log(minimumAddedInteger([]int{9, 4, 3, 9, 4}, []int{7, 8, 8}))                           // 4
	t.Log(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10}))
	t.Log(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))
}
