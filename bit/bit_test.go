package bit

import (
	"fmt"
	"testing"
)

// 区域和检索 - 数组不可变
//func TestNumArray(t *testing.T) {
//	// -2 -2  1 -4 -2 -3
//	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
//	// 1
//	fmt.Println(numArray.SumRange(0,2))
//	// -1
//	fmt.Println(numArray.SumRange(2,5))
//	// -3
//	fmt.Println(numArray.SumRange(0,5))
//}
//
//type NumArray struct {
//	nums []int
//}
//
//
//func Constructor(nums []int) NumArray {
//	temp := make([]int,len(nums) + 1)
//	for i := 0; i < len(nums); i++ {
//		temp[i + 1] = nums[i] + temp[i]
//	}
//	return NumArray{
//		nums: temp,
//	}
//}
//
//
//func (this *NumArray) SumRange(left int, right int) int {
//	return this.nums[right + 1] - this.nums[left]
//}

// 区域和检索 - 数组可修改
func TestNumArray1(t *testing.T) {
	// -2 -2  3 -4  2  1
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	// 1
	fmt.Println(numArray.SumRange(0, 2))
	// -1
	fmt.Println(numArray.SumRange(2, 5))
	// -3
	fmt.Println(numArray.SumRange(0, 5))

	numArray = Constructor([]int{1, 3, 5})
	// 1
	fmt.Println(numArray.SumRange(0, 2))
	numArray.Update(1, 2)
	// -3
	fmt.Println(numArray.SumRange(0, 2))
}

func lowbit(x int) int {
	return x & -x
}

// 查询前缀和
func (this *NumArray) query(x int) int {
	res := 0
	// 注：等于0 会导致死循环
	for i := x; i > 0; i -= lowbit(i) {
		res += this.tree[i]
	}
	return res
}

// 在树状数组 k 位置增加 v，更新该节点与该节点所有的父节点
func (this *NumArray) add(k, v int) {
	for i := k; i < len(this.tree); i += lowbit(i) {
		this.tree[i] += v
	}
}

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	temp := make([]int, l+1)
	// 构造树状数组
	for i := 1; i <= l; i++ {
		temp[i] += nums[i-1]
		// 父节点
		j := i + lowbit(i)
		if j <= l {
			temp[j] += temp[i]
		}
	}
	return NumArray{nums: nums, tree: temp}
}

func (this *NumArray) Update(index int, val int) {
	// 原有的值是 nums[index]，要使得修改为 val，所以需要增加 val - nums[index]
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right+1) - this.query(left)
}

// 算法题
func TestName(t *testing.T) {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
	fmt.Println(truncateSentence("chopper is not a tanuki", 5))
}

func truncateSentence(s string, k int) string {
	index := 0
	for i, i2 := range s {
		if i2 == ' ' {
			k--
			index = i
		}
		if k == 0 {
			break
		}
	}
	if k != 0 {
		return s
	}
	return s[:index]
}
