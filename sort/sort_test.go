package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//nums := []int{3, 4, 2, 1, 5, 7, 6}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	bubbleSort2(nums)
	fmt.Println(nums)
}

// 冒泡排序
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

// 冒泡排序
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

// 冒泡排序
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
	nums := []int{3, 4, 2, 1, 5, 7, 6}
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	insertionSort(nums)
	fmt.Println(nums)
}

// 插入排序
func insertionSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		temp := nums[i]
		index := i - 1
		for index >= 0 && nums[index] >= temp {
			nums[index+1] = nums[index]
			index--
		}
		nums[index+1] = temp
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
	return merge(left, right)
}

func merge(left, right []int) (res []int) {
	leftL := len(left)
	rightL := len(right)
	indexL, indexR := 0, 0
	for indexL < leftL && indexR < rightL {
		if left[indexL] < right[indexR] {
			res = append(res, left[indexL])
			indexL++
		} else {
			res = append(res, right[indexR])
			indexR++
		}
	}
	res = append(res, left[indexL:]...)
	res = append(res, right[indexR:]...)
	return
}

// 排序链表
func TestSortList(t *testing.T) {
	root := &ListNode{
		4,
		&ListNode{
			2,
			&ListNode{
				1,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	res := sortList(root)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	res := mergeList(left, right)
	return res
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	var head = &ListNode{
		0,
		nil,
	}
	temp := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			temp.Next = left
			left = left.Next
		} else {
			temp.Next = right
			right = right.Next
		}
		temp = temp.Next

	}
	if left != nil {
		temp.Next = left
	}
	if right != nil {
		temp.Next = right
	}
	return head.Next
}

// 剑指 Offer 51. 数组中的逆序对
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
// 示例 1:
//
// 输入: [7,5,6,4]
// 输出: 5
func TestReversePairs(t *testing.T) {
	//nums := []int{4,5,6,7}
	//nums := []int{7,5,6,4}
	nums := []int{1, 3, 2, 3, 1}
	res := reversePairs(nums)
	fmt.Println(res)
}

func reversePairs(nums []int) int {
	_, res := mergeSortPairs(nums)
	return res
}

func mergeSortPairs(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}
	mid := len(nums) / 2
	left, countL := mergeSortPairs(nums[:mid])
	right, countR := mergeSortPairs(nums[mid:])
	res, count := mergePairs(left, right)
	return res, countL + countR + count
}

func mergePairs(left, right []int) (res []int, count int) {
	leftL := len(left)
	rightL := len(right)
	indexL, indexR := 0, 0
	for indexL < leftL && indexR < rightL {
		if left[indexL] <= right[indexR] {
			res = append(res, left[indexL])
			indexL++
		} else {
			count += leftL - indexL
			res = append(res, right[indexR])
			indexR++
		}
	}
	res = append(res, left[indexL:]...)
	res = append(res, right[indexR:]...)
	return
}

// 快速排序
func TestQuickSort(t *testing.T) {
	nums := []int{3, 4, 2, 1, 5, 7, 6}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, start int, end int) {
	if start < end {
		index := quick(nums, start, end)
		quickSort(nums, start, index-1)
		quickSort(nums, index+1, end)
	}
}

func quick(nums []int, start int, end int) int {
	num := nums[start]
	index := start + 1
	for i := index; i <= end; i++ {
		if nums[i] < num {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index-1], nums[start] = nums[start], nums[index-1]
	return index - 1
}

// 优化
func quick1(nums []int, start int, end int) int {
	num := nums[start]
	index := start + 1
	for index <= end {
		for index <= end && nums[index] < num {
			index++
		}
		for index <= end && nums[end] > num {
			end--
		}
		if index <= end {
			nums[index], nums[end] = nums[end], nums[index]
			index++
		}
	}
	nums[index-1], nums[start] = nums[start], nums[index-1]
	return index - 1
}

// 优化
func quick2(nums []int, start int, end int) int {
	randomIndex := rand.Intn(end-start+1) + start                   // 随机选择基准元素的索引
	nums[start], nums[randomIndex] = nums[randomIndex], nums[start] // 将随机选择的元素交换到数组起始位置
	num := nums[start]
	index := start + 1
	for index <= end {
		for index <= end && nums[index] < num {
			index++
		}
		for index <= end && nums[end] > num {
			end--
		}
		if index <= end {
			nums[index], nums[end] = nums[end], nums[index]
			index++
		}
	}
	nums[index-1], nums[start] = nums[start], nums[index-1]
	return index - 1
}

// 堆排序
func TestHeapSort(t *testing.T) {
	nums := []int{3, 4, 2, 1, 5, 7, 6}
	heapSort(nums)
	fmt.Println(nums)
}

func heapSort(nums []int) {
	l := len(nums)
	buildMaxHeap(nums, l)
	for i := l - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		l -= 1
		heapify(nums, 0, l)
	}
}

func buildMaxHeap(nums []int, l int) {
	for i := l / 2; i >= 0; i-- {
		heapify(nums, i, l)
	}
}

func heapify(nums []int, i int, l int) {
	left := 2*i + 1
	right := 2*i + 2
	index := i
	if left < l && nums[left] > nums[index] {
		index = left
	}
	if right < l && nums[right] > nums[index] {
		index = right
	}
	if index != i {
		nums[index], nums[i] = nums[i], nums[index]
		heapify(nums, index, l)
	}
}
