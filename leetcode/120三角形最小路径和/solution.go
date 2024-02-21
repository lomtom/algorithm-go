package leetcode

import (
	"fmt"
	"testing"
)

// 2
// 3 4
// 6 5 7
// 4 1 8 3
// 执行耗时:4 ms,击败了75.62% 的Go用户
// 内存消耗:2.8 MB,击败了93.59% 的Go用户
func minimumTotal1(triangle [][]int) int {
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	for index1 := range triangle {
		if index1 == 0 {
			continue
		}
		for index2 := range triangle[index1] {
			if index2 == 0 { // 第一个元素
				triangle[index1][index2] += triangle[index1-1][index2]
			} else if index2 == len(triangle[index1])-1 { // 最后一个元素
				triangle[index1][index2] += triangle[index1-1][index2-1]
			} else {
				triangle[index1][index2] += min(triangle[index1-1][index2-1], triangle[index1-1][index2])
			}
		}
	}
	var number = triangle[len(triangle)-1][0]
	for index := range triangle[len(triangle)-1] {
		number = min(number, triangle[len(triangle)-1][index])
	}
	return number
}

// 2
// 3 4
// 6 5 7
// 4 1 8 3
// 执行耗时:4 ms,击败了75.62% 的Go用户
// 内存消耗:2.8 MB,击败了100.00% 的Go用户
func minimumTotal(triangle [][]int) int {
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	for index1 := len(triangle) - 2; index1 >= 0; index1-- {
		for index2 := range triangle[index1] {
			triangle[index1][index2] += min(triangle[index1+1][index2], triangle[index1+1][index2+1])
		}
	}
	return triangle[0][0]
}

func TestMinimumTotal(t *testing.T) {
	var triangle [][]int
	//triangle = make([][]int, 1)
	//triangle[0] = []int{-10}
	triangle = make([][]int, 4)
	triangle[0] = []int{2}
	triangle[1] = []int{3, 4}
	triangle[2] = []int{6, 5, 7}
	triangle[3] = []int{4, 1, 8, 3}
	fmt.Println(minimumTotal(triangle))
}
