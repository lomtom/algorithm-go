package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(dominantIndex([]int{100, 1, 50})) // 0
	t.Log(dominantIndex([]int{3, 6, 1, 0})) // 1
	t.Log(dominantIndex([]int{1, 2, 3, 4})) // -1
	t.Log(dominantIndex([]int{1}))          // 0
	t.Log(dominantIndex([]int{0}))          // 0
	t.Log(dominantIndex([]int{0, 6}))       // 1
	t.Log(dominantIndex([]int{10, 6}))      // -1
	t.Log(dominantIndex([]int{0, 100}))     // 1
	t.Log(dominantIndex([]int{0, 0, 2, 1})) // 2
	t.Log(dominantIndex([]int{0, 50, 100})) // 2
}
