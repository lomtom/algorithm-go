package dfs

import (
	"fmt"
	"testing"
)

// 全排列
func TestPermute(t *testing.T) {
	res := permute([]int{5, 4, 6, 2})
	fmt.Println(res)
	res = permute([]int{1, 2, 3})
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	flag := make([]bool, len(nums), len(nums))
	list := make([]int, 0)
	permuteDfs(nums, list, &res, len(nums), &flag)
	return res
}

func permuteDfs(nums, list []int, res *[][]int, l int, flag *[]bool) {
	if len(list) == l {
		*res = append(*res, list)
		return
	}
	for i := 0; i < l; i++ {
		if (*flag)[i] {
			continue
		}
		list = append(list, (nums)[i])
		(*flag)[i] = true
		permuteDfs(nums, list, res, l, flag)
		list = (list)[:len(list)-1]
		(*flag)[i] = false
	}
}

func TestName(t *testing.T) {
	res := make([][]int, 0)
	flag := make([]bool, 4, 4)
	list := make([]int, 0)
	fn(&res, list, flag)
	fmt.Println(res)
}

func fn(res *[][]int, list []int, flag []bool) {
	if len(list) == 4 {
		*res = append(*res, list)
		return
	}
	for i := 0; i < 4; i++ {
		if !flag[i] {
			flag[i] = true
			list = append(list, i)
			fn(res, list, flag)
			list = list[:len(list)-1]
			flag[i] = false
		}
	}
}
