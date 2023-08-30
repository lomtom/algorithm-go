---
title: "排序"
slug: "sort"
---
排序算法有很多算法，例如

简单，也是用的最多的有
- [冒泡排序](#冒泡排序)
- [选择排序](#选择排序)
- [插入排序](#插入排序)

稍微复杂的有：
- [归并排序](#归并排序)
- [快速排序](#快速排序)
- 堆排序

其他还有：
- [希尔排序](#希尔排序)
- [计数排序](#计数排序)
- 桶排序
- 等等

排序算法对于我们来说是经常会使用到的一种算法，也是数据结构中最基本的算法，也许你会说我直接用官方语言提供的方法或者api不就好了吗？

那确实也没错，用官方提供的方法确实很方便快捷，但是作为对知识好求的我们，难道不想知道其所以然吗？

更何况，排序算法也是日后面试必问的问题之一呀，而且考证啥的也会涉及到排序算法的。

**所以，不会排序的赶快学起来吧。**


## 冒泡排序
### 冒泡排序


冒泡排序（Bubble Sort）是一种简单的排序算法，它重复地遍历要排序的列表，比较相邻的两个元素，并将它们交换位置，如果它们的顺序错误。这个过程持续进行，直到整个列表排序完成。冒泡排序算法的名字就来源于元素不断像气泡一样冒到正确的位置。

冒泡排序的基本思想是每次比较相邻的两个元素，如果它们的顺序错误就交换它们。这样，经过一轮的遍历，最大的元素就会“冒泡”到最右边的位置。然后，对剩下的未排序元素进行相同的操作，直到整个数组排序完成。

以下是冒泡排序的基本步骤：

- 从列表的第一个元素开始，依次比较相邻的两个元素，如果顺序错误就交换它们的位置，将较大的元素“冒泡”到右边。
- 继续对剩下的元素进行相同的操作，直到所有元素都被比较过一次。
- 重复上述步骤，但是每次循环减少一个元素，因为每一轮都会将当前轮中最大的元素“沉”到最后。

冒泡排序是一种相对简单但效率较低的排序算法，因为在每一轮遍历中，都需要进行多次比较和交换。在最坏的情况下，需要进行 n-1 轮遍历，每轮遍历需要进行 n-i 次比较和交换，其中 n 是列表中的元素数量。

![](../img/sort/bubbleSort.gif)

**由于其复杂度，一般我们避免使用冒泡排序**

首先，先来一个最简单版的冒泡排序
```go
// 冒泡算排序
func bubbleSort(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := 0; j < len - 1; j++ {
			if nums[j] >= nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
}
```

### 优化1

**分析：**
如果你稍微分析一下代码，你就会知道这个代码其实做了一些无用的比较，例如，排好序的元素，还是发生了比较。

那么，怎么进行优化呢？


**方案：**
只需要将排好序的元素不再进行比较不就行了，每次循环排好一个元素，所以每次循环比较时，减少一次比较，也就是i次。

那么稍微优化一下你就可以得到
```go
// 冒泡算排序
func bubbleSort1(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := 0; j < len - 1 - i; j++ {
			if nums[j] >= nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
}
```

### 优化2


**分析：**
到这里，觉得差不多了，很多同学到这里也就终止了，其实，在这里还是有一个点，也是经常会被忽略的点。

1. 我们知道，冒泡排序最坏的情况复杂度为O(n^2)，也就是每个元素，每次比较都进行交换，也就是原数组顺序与我们得到的顺序相反
2. 而最好的情况复杂度同样是为O(n^2)，也就是说尽管原数组顺序与我们期望的顺序一致，仍然需要比较n^2

那么怎么在情况二的条件下降低复杂度呢？


**方案：**
我们可以使用临时标记，如果一个循环的比较未发生交换时，就代表数组顺序是我们期望的，此时就可以跳出循环，从而降低复杂度，也就是O(n)。

```go
// 冒泡算排序
func bubbleSort2(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		flag := false
		for j := 0; j < len - 1 - i; j++ {
			if nums[j] >= nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
```
## 选择排序

选择排序（Selection Sort）是一种简单直观的排序算法，它的基本思想是将待排序的序列分为已排序区和未排序区，每次从未排序区中选取最小（或最大）的元素，将其与已排序区的末尾元素交换位置，从而逐步构建有序序列。

选择排序的主要步骤如下：
- 在未排序区中找到最小（或最大）的元素。
- 将该元素与未排序区的第一个元素交换位置，将其纳入已排序区。
- 重复上述步骤，每次在未排序区中找到最小（或最大）的元素，将其交换到已排序区的末尾。

![](../img/sort/selectionSort.gif)
```go
// 选择排序
func selectionSort(nums []int)  {
	l := len(nums)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			nums[i],nums[min] = nums[min],nums[i]
		}
	}
}
```

选择排序相对于冒泡排序，并不需要进行多余的交换，所以性能上优于冒泡排序（仅限于数组规模较小时）。

## 插入排序
### 插入排序

插入排序（Insertion Sort）是一种简单且常用的排序算法，它的基本思想是将待排序的序列分成已排序区和未排序区，每次从未排序区选取一个元素，将其插入到已排序区的合适位置，从而逐步构建有序序列。

插入排序的主要步骤如下：
- 将第一个元素视为已排序区，将剩余元素视为未排序区。
- 从未排序区选择一个元素，将其插入到已排序区的合适位置，使得插入后仍然保持有序。
- 重复上述步骤，直到未排序区中的所有元素都被插入到已排序区。

![](../img/sort/insertionSort.gif)


```go
// 插入排序
func insertionSort(nums []int)  {
    l := len(nums)
    for i := 0; i < l; i++ {
        temp := nums[i]
        index := i - 1
        for  index >= 0 && nums[index] >= temp {
            nums[index + 1] = nums[index]
            index--
        }
        nums[index + 1] = temp
    }
}
```
### 例题
[对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/)

## 归并排序
### 归并排序
归并排序（Merge Sort）是一种基于[分治策略](dac)的排序算法，它将一个未排序的数组划分成两个子数组，然后递归地对子数组进行排序，最后将排好序的子数组合并以得到完整的排序数组。归并排序的核心思想是将问题分解成更小的子问题，解决子问题，然后合并子问题的解来解决原始问题。

![](../img/sort/mergeSort.gif)

核心过程：
1. 分：将未排序的数组分解成两个子数组，通常通过将数组从中间切分来实现。分解过程一直持续（采用递归），直到每个子数组都只包含一个元素为止，因为单个元素可以被认为是有序的。
2. 治：递归地对分解得到的子数组进行排序。这是通过不断地将问题分解为更小的子问题，然后逐个解决这些子问题，直到达到基本情况（单个元素）为止。
3. 合：合并已排序的子数组，以获得一个完整的、有序的数组。合并过程涉及比较两个子数组的元素，并按顺序将它们合并到一个新的数组中。

归并排序的关键在于合并步骤，它要求在合并两个有序的子数组时保持它们的有序性。这可以通过比较两个子数组的第一个元素，选择较小的元素放入结果数组中，然后逐个向后移动指针，继续比较和合并元素。

![](../img/sort/mergeSort.webp)

```go
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
```

### 例题
1. [排序数组](https://leetcode-cn.com/problems/sort-an-array/)

2. [排序链表](https://leetcode-cn.com/problems/sort-list/)

3. [剑指 Offer 51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

4. [翻转对](https://leetcode-cn.com/problems/reverse-pairs/)

5. [计算右侧小于当前元素的个数](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/)

### 题解
1、排序数组：略

2、排序链表

难点在于怎么进行拆分，链表没办法像数组那样直接获取中间节点，那么我们可以采用快慢指针找到中间节点。

```go
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
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
	res := mergeList(left,right)
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
		} else{
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

```

3、剑指 Offer 51. 数组中的逆序对

这题很难想到会用归并排序的思想来解题，然而这道题确实归并排序的经典题。

拆分后，只需对前后两个数组进行对比即可，left数组的第indexL个对比right的的第indexR个，如果符合条件，逆序对加leftL - indexL即可。

```go

func reversePairs(nums []int) int {
	_,res :=mergeSortPairs(nums)
	return res
}

func mergeSortPairs(nums []int) ([]int,int) {
	if len(nums) <= 1 {
		return nums,0
	}
	mid := len(nums)/2
	left,countL := mergeSortPairs(nums[:mid])
	right,countR := mergeSortPairs(nums[mid:])
	res,count := mergePairs(left,right)
	return res,countL + countR + count
}

func mergePairs(left,right []int) (res []int,count int) {
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
```


4、翻转对

```go
func reversePairs(nums []int) int {
	_, res := mergeSort(nums)
	return res
}

func mergeSort(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}
	mid := len(nums) / 2
	left, countL := mergeSort(nums[:mid])
	right, countR := mergeSort(nums[mid:])
	res, count := merge(left, right)
	return res, count + countL + countR
}

func merge(left, right []int) (res []int, count int) {
	leftL := len(left)
	rightL := len(right)
	indexL, indexR := 0, 0
	j := 0
	for _, v := range left {
		for j < rightL && v > 2*right[j] {
			j++
		}
		count += j
	}
	for indexL < leftL && indexR < rightL {
		if left[indexL] <= right[indexR] {
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
```

## 快速排序

快速排序（Quick Sort）是一种高效的、基于分治思想的排序算法，由英国计算机科学家 Tony Hoare 于1960年提出。它在平均情况下具有很好的性能，通常被认为是最快的排序算法之一。

快速排序的基本思想如下：

1. 选择基准元素： 从待排序的数组中选择一个基准元素，通常是数组中的第一个元素或者随机选择。基准元素的选择会影响快速排序的性能。

2. 分割操作： 将数组中的元素分成两部分，一部分小于基准元素，另一部分大于基准元素。这个过程称为分割操作（Partition）。

3. 递归排序： 对分割后的两个部分分别进行递归排序，即对小于基准的部分和大于基准的部分分别应用快速排序。

4. 合并： 无需合并步骤，因为快速排序在原地进行排序，即直接在原始数组上进行元素交换，而不需要额外的内存空间。

整个过程不断地递归进行，直到每个子数组只包含一个元素或没有元素为止，排序完成。


![](../img/sort/quickSort.gif)

```go
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
	for i := start; i <= end; i++ {
		if nums[i] < num {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index-1], nums[start] = nums[start], nums[index-1]
	return index - 1
}
```

### 优化1
上面的实现中，每次分割都需要遍历整个数组，这样会导致时间复杂度为 O(n^2)。可以通过优化分割操作来减少时间复杂度。
```go
// 优化
func quick1(nums []int, start int, end int) int {
	num := nums[start]
	index := start + 1
	for index <= end {
		for index <= end && nums[index] > num {
			index++
		}
		for index <= end && nums[end] < num {
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
```

### 优化2
为了在极端情况下，比如数组已经有序，或者逆序，导致每次分割只能减少一个元素，这种情况下，快速排序的时间复杂度会退化为 O(n^2)。为了避免这种情况，可以在每次分割时，随机选择一个元素作为基准元素。
```go
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
```
## 希尔排序

希尔排序（Shell Sort）是一种排序算法，也被称为缩小增量排序。它是插入排序的一种高效改进版本，旨在克服插入排序在处理大规模数据时的低效率问题。希尔排序通过将整个待排序的数组分成一些小的子数组，针对每个子数组进行插入排序，然后逐步减小子数组的规模，最终实现整个数组的排序。

希尔排序的核心思想是先对间隔较大的元素进行排序，使得数组中的每个元素都能离其正确位置更近一些。随着算法的进行，间隔逐渐缩小，最终变为1，此时数组已经基本有序，只需要进行少量的插入排序即可。

希尔排序的步骤如下：

- 选择一个增量序列（间隔序列），通常是n/2，然后依次减半，直到增量为1。
- 对每个增量，将数组分成若干个子数组，每个子数组中的元素相隔增量个位置。
- 对每个子数组进行插入排序，使得子数组中的元素基本有序。
- 重复步骤2和3，直到增量为1时，对整个数组进行最后一次插入排序。

希尔排序的优势在于它在初始阶段就可以将大部分数据移动到正确的位置，从而减少了插入排序中大量的元素交换次数，提高了排序的效率。然而，希尔排序的性能依赖于增量序列的选择，不同的增量序列可能导致不同的性能表现。

尽管希尔排序相对于其他复杂度更低的排序算法（如快速排序、归并排序）来说性能较差，但在某些特定情况下（例如对中小规模数据进行排序），它的表现仍然是不错的。希尔排序也是一种原地排序算法，不需要额外的内存空间。

```go
func shellSort(arr []int) {
	n := len(arr)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] < temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2
	}
}
```

## 计数排序

计数排序（Counting Sort）是一种非比较性质的排序算法，适用于一定范围内的整数排序。它利用输入的元素与范围内的整数建立映射，通过统计每个整数出现的次数来对输入元素进行排序。计数排序的核心思想是将输入元素分布到不同的“桶”中，然后根据桶中元素的顺序重建排序结果。

以下是计数排序的基本步骤：

- 寻找输入的范围： 首先确定输入元素的范围，找出最小和最大的整数值，以确定需要创建多大的计数数组。

- 创建计数数组： 创建一个长度为范围大小的计数数组，用于统计每个整数出现的次数。计数数组的索引对应于整数值，值为该整数出现的次数。

- 累加计数数组： 将计数数组中的值进行累加，以表示小于或等于该索引的整数有多少个。这一步将帮助确定每个整数在排序结果中的位置。

- 排序操作： 遍历输入数组，将每个元素放入到计数数组中对应的位置，并将计数数组中对应的值减1。这一步操作会将输入元素按照计数数组的顺序放入正确的位置，从而实现排序。

![](../img/sort/countingSort.gif)

计数排序的优点在于它对于一定范围内的整数排序具有很高的性能，因为它不需要比较操作。然而，计数排序的主要限制在于它需要额外的存储空间来存储计数数组，因此在输入范围较大时可能会消耗大量的内存。

计数排序的时间复杂度是 O(n + k)，其中 n 是输入元素的个数，k 是整数范围的大小。当输入范围不是很大时，计数排序可以是一种高效的排序方法。

需要注意的是，计数排序只适用于整数排序，而且在输入范围很大时，其效率可能会受到限制。

```go

func countingSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	// 寻找输入的范围
	var max = nums[0]
	var min = nums[0]
	for index := range nums {
		if nums[index] > max {
			max = nums[index]
		}
		if nums[index] < min {
			min = nums[index]
		}
	}
	if min < 0 {
		min = -min
	}
	// 创建计数数组
	counts := make([]int, max+min+1)
	// 累加计数数组
	for index := range nums {
		counts[nums[index]+min]++
	}
	// 排序操作
	index := 0
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			nums[index] = i - min
			index++
		}
	}
}
```
