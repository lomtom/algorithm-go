package leetcode

import (
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	t.Log(nextGreaterElements([]int{1, 2, 1}))
	t.Log(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
