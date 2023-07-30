package leetcode

import (
	"fmt"
	"testing"
)

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.7 MB,击败了27.27% 的Go用户
func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int][]int, 0)
	for index := range pieces {
		m[pieces[index][0]] = pieces[index]
	}
	index := 0
	for index < len(arr) {
		if v, ok := m[arr[index]]; !ok {
			return false
		} else {
			for i := range v {
				if v[i] == arr[index] {
					index++
				} else {
					return false
				}
			}
		}
	}
	return true
}

func TestCanFormArray(t *testing.T) {
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(canFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}
