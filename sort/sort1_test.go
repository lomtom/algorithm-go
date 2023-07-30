package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	nums := []int{4, 3, 2, 1, 5, 6, 7, 0}
	//BubbleSort(nums)
	//SelectSort(nums)
	//InsertSort(nums)
	fmt.Println(MergeSort(nums))
	fmt.Println(nums)
}

func BubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		flag := false
		for j := 0; j < l-i-1; j++ {
			if nums[j] < nums[j+1] {
				flag = true
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if !flag {
			break
		}
	}
}

func SelectSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		index := i
		num := nums[i]
		for j := i; j < l; j++ {
			if nums[j] >= num {
				num = nums[j]
				index = j
			}
		}
		if i != index {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}

func InsertSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		index := i - 1
		num := nums[i]
		for index >= 0 && num > nums[index] {
			nums[index+1] = nums[index]
			index--

		}
		nums[index+1] = num
	}
}

func MergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	mid := l / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	lLeft := len(left)
	lRight := len(right)
	indexLeft, indexRight := 0, 0
	res := make([]int, 0)
	for indexLeft < lLeft && indexRight < lRight {
		if left[indexLeft] < right[indexRight] {
			res = append(res, left[indexLeft])
			indexLeft++
		} else {
			res = append(res, right[indexRight])
			indexRight++
		}
	}
	if indexLeft < lLeft {
		res = append(res, left[indexLeft:]...)
	}
	if indexRight < lRight {
		res = append(res, right[indexRight:]...)
	}
	return res
}

func TestName(t *testing.T) {
	nums := []int{15,2,8,13}
	fmt.Println(nextGreaterElement(nums))
}

func nextGreaterElement(nums []int) []int {
	l := len(nums)
	result := make([]int, l)
	stack := make([]int,0)
	for i := l - 1;i >= 0;i-- {
		for len(stack) != 0 && nums[i] >= stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) == 0 {
			result[i] = -1
		}else{
			result[i] = stack[len(stack) - 1]
		}
		stack = append(stack,nums[i])
	}
	return result
}
