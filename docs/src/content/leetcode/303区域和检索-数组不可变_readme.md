---
title: 303区域和检索-数组不可变
categories:
  - 简单
tags:
  - 数组
number: 303
---

**题目难度：**[区域和检索 - 数组不可变](https://leetcode.cn/problems/range-sum-query-immutable/) 简单

**题目描述：**

> 给定一个整数数组  nums，处理以下类型的多个查询:
> 
> 
> 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
> 
> 实现 NumArray 类：
> 
> NumArray(int[] nums) 使用数组 nums 初始化对象
> 
> int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )


**测试用例：**

> 示例 1:
>
> 输入：
> 
> ["NumArray", "sumRange", "sumRange", "sumRange"]
> 
> [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
> 
> 输出：
> 
> [null, 1, -1, -3]
> 
> 解释：
> 
> NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
> 
> numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
> 
> numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
> 
> numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))


**限制及提示：**
> - 1 <= nums.length <= 104
> - -105 <= nums[i] <= 105
> - 0 <= i <= j < nums.length
> - 最多调用 104 次 sumRange 方法


---
**解题分析及思路：**


该题直接用暴力每次计算left到right之间值的和即可，但是也可以使用[树状数组](../pages/bit)来进行解答。

首先建立一棵二叉索引树。

```go
func Constructor(nums []int) NumArray {
	l := len(nums)
	ints := make([]int, l+1)
	for index := 1; index <= l; index++ {
		index1 := index
		for index1 <= l {
			ints[index1] += nums[index-1]
			index1 += lowBit(index1)
		}
	}
	return NumArray{
		ints,
	}
}
```

然后每次计算的两个区间的值的时候，计算sum(right + 1) - sum(left)即可。
```go
func sum(ints []int, index int) (ans int) {
	for index > 0 {
		ans += ints[index]
		index -= lowBit(index)
	}
	return ans
}
```



[源代码](https://github.com/lomtom/algorithm-go/blob/main/leetcode/303/303区域和检索-数组不可变_test.go)

**复杂度：**
- 时间复杂度：O(logn)
- 空间复杂度：O(n)

**执行结果：**

- 执行用时：24 ms, 在所有 Go 提交中击败了82.00%的用户
- 内存消耗：9.8 MB, 在所有 Go 提交中击败了21.91%的用户
