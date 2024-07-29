package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(calPoints([]string{"5", "2", "C", "D", "+"}))
	t.Log(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
