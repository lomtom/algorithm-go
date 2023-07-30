package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:1.9 MB,击败了100.00% 的Go用户
func canPartitionKSubsets(nums []int, k int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	if total%k != 0 {
		return false
	}
	var avg = total / k
	sort.Ints(nums)
	var flag = make([]bool, len(nums))
	var dfs func(index int, sum int, count int) bool
	dfs = func(index int, sum int, count int) bool {
		if k == count {
			return true
		}
		if sum == avg {
			return dfs(len(nums)-1, 0, count+1)
		}
		if index == -1 {
			return false
		}
		for i := index; i >= 0; i-- {
			if flag[i] || sum+nums[i] > avg {
				continue
			}
			flag[i] = true
			if dfs(i-1, sum+nums[i], count) {
				return true
			}
			flag[i] = false
			if sum == 0 {
				return false
			}
		}
		return false
	}
	return dfs(len(nums)-1, 0, 0)
}

func TestCanPartitionKSubsets(t *testing.T) {
	// false
	fmt.Println(canPartitionKSubsets([]int{4, 2, 9, 8, 1, 1, 5, 9, 4, 3, 5, 6, 3, 5, 7}, 9))
	// false
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	// false
	fmt.Println(canPartitionKSubsets([]int{1, 1, 2, 4}, 4))
	// true
	fmt.Println(canPartitionKSubsets([]int{4, 5, 9, 3, 10, 2, 10, 7, 10, 8, 5, 9, 4, 6, 4, 9}, 5))
	// true
	fmt.Println(canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2}, 4))
	// true
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	// false
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	// true
	fmt.Println(canPartitionKSubsets([]int{1, 3, 4}, 2))
}
