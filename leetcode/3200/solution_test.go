package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxHeightOfTriangle(2, 4))  // 3
	t.Log(maxHeightOfTriangle(2, 1))  // 2
	t.Log(maxHeightOfTriangle(1, 1))  // 1
	t.Log(maxHeightOfTriangle(10, 1)) // 2
}
