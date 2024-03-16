package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(canFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
}
