package leetcode

import (
	"fmt"
	"testing"
)

func maxAscendingSum(nums []int) int {
	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	sum, tempSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tempSum += nums[i]
			continue
		}
		sum = max(sum, tempSum)
		tempSum = nums[i]
	}
	return max(sum, tempSum)
}

func TestMaxAscendingSum(t *testing.T) {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}
