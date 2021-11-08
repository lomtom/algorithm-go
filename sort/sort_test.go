package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{3, 4, 2, 1, 5, 7, 6}
	res := mergeSort(nums)
	fmt.Println(res)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	res := merge(left, right)
	return res
}

func merge(left []int, right []int) (res []int) {
	leftLen := len(left)
	rightLen := len(right)
	i := 0
	j := 0
	for i < leftLen && j < rightLen {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

func TestQuickSort(t *testing.T) {
	nums := []int{3, 4, 2, 1, 5, 7, 6}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, start, end int) {
	if start < end {
		index := quick(nums, start, end)
		quickSort(nums, start, index-1)
		quickSort(nums, index+1, end)
	}
}

func quick(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
