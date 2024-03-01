package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(validPartition([]int{1, 2, 3}))
	t.Log(validPartition([]int{993335, 993336, 993337, 993338, 993339, 993340, 993341}))
	t.Log(validPartition([]int{1, 1, 1}))
}
