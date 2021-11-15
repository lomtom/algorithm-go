package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//nums := []int{3, 4, 2, 1, 5, 7, 6}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	bubbleSort2(nums)
	fmt.Println(nums)
}

// 冒泡算排序
func bubbleSort(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 冒泡算排序
func bubbleSort1(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1-i; j++ {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 冒泡算排序
func bubbleSort2(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		flag := false
		for j := 0; j < len-1-i; j++ {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func TestSelectionSort(t *testing.T) {
	//nums := []int{3, 4, 2, 1, 5, 7, 6}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	selectionSort(nums)
	fmt.Println(nums)
}

// 选择排序
func selectionSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}

func TestInsertionSort(t *testing.T) {
	//nums := []int{3, 4, 2, 1, 5, 7, 6}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	insertionSort(nums)
	fmt.Println(nums)
}

// 插入排序
func insertionSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		temp := nums[i]
		index := i
		for index > 0 && nums[index] >= temp {
			nums[index] = nums[index-1]
		}
		nums[index] = temp
	}
}

// 归并排序
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
