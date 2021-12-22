# 排序算法

排序算法有很多算法，例如

简单，也是用的最多的有
- 冒泡排序
- 选择排序
- 插入排序

稍微复杂的有：
- 归并排序
- 快速排序
- 堆排序

其他还有：
- 希尔排序
- 计数排序
- 桶排序
- 等等

排序算法对于我们来说是经常会使用到的一种算法，也是数据结构中最基本的算法，也许你会说我直接用官方语言提供的方法或者api不就好了吗？

那确实也没错，用官方提供的方法确实很方便快捷，但是作为对知识好求的我们，难道不想知道其所以然吗？

更何况，排序算法也是日后面试必问的问题之一呀，而且考证啥的也会涉及到排序算法的。

**所以，不会排序的赶快学起来吧。**


## 冒泡排序
### 冒泡排序
冒泡排序是与插入排序拥有相等的渐近时间复杂度，但是两种算法在需要的交换次数却很大地不同。


冒泡排序算法的运作如下：
1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
3. 针对所有的元素重复以上的步骤，除了最后一个。
4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

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
选择排序的核心就是以第一个未排序的元素作为基准，在该元素的之后的元素中找到一个最小的元素（比该元素小），与其交换位置

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
插入排序也比较简单，其核心思想就是找到第一个未排序的元素，插入到该元素前排序好的元素当中去（与前面的元素进行比较，如果大于当前元素，将其往后挪一个，否则，插入到当前位置）。

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
归并排序相对于前几种排序来说，复杂多了，当然他理解起来，也并不复杂。

归并排序的核心算法是[分治法](dac.md)。

![](../img/sort/mergeSort.gif)

核心过程：
1. 分：使用递归的方法，将数组每次切成两部分，直到数组的大小为1，即不可再切割，此时，每个子数组都是有序的。
2. 治：从底层合并两个排好序的数组，因为是有序的，所以从数组的第一个元素开始比较即可，小的放入一个新的数组。

上动图并不能直观的的表达归并排序的核心思想，可以参考下图。

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



