package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxValue(1, 0, 24)) // 3 ： 2 3 2 1 1 1 / 2 2 2 1 1 1 / 5 5 5 4 3 2
	t.Log(maxValue(6, 1, 10)) // 3 ： 2 3 2 1 1 1 / 2 2 2 1 1 1 / 5 5 5 4 3 2
	t.Log(maxValue(4, 2, 6))
}
