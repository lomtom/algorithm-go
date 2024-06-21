---
title: 气温变化趋势
categories:
  - 简单
tags: 
  - 数组
slug: 6CE719
number: LCP61
---

## 题目描述：

力扣城计划在两地设立「力扣嘉年华」的分会场，气象小组正在分析两地区的气温变化趋势，对于第 `i ~ (i+1)` 天的气温变化趋势，将根据以下规则判断：
- 若第 `i+1` 天的气温 **高于** 第 `i` 天，为 **上升** 趋势
- 若第 `i+1` 天的气温 **等于** 第 `i` 天，为 **平稳** 趋势
- 若第 `i+1` 天的气温 **低于** 第 `i` 天，为 **下降** 趋势

已知 `temperatureA[i]` 和 `temperatureB[i]` 分别表示第 `i` 天两地区的气温。
组委会希望找到一段天数尽可能多，且两地气温变化趋势相同的时间举办嘉年华活动。请分析并返回两地气温变化趋势**相同的最大连续天数**。
> 即最大的 `n`，使得第 `i~i+n` 天之间，两地气温变化趋势相同

**示例 1：**
>输入：
>`temperatureA = [21,18,18,18,31]`
>`temperatureB = [34,32,16,16,17]`
>
>输出：`2`
>
>解释：如下表所示， 第 `2～4` 天两地气温变化趋势相同，且持续时间最长，因此返回 `4-2=2`

![image.png](/img/leetcode/LCP61气温变化趋势/1663902654-hlrSvs-image.png){:width=1000px}


**示例 2：**
>输入：
>`temperatureA = [5,10,16,-6,15,11,3]`
>`temperatureB = [16,22,23,23,25,3,-16]`
>
>输出：`3`

**提示：**
- `2 <= temperatureA.length == temperatureB.length <= 1000`
- `-20 <= temperatureA[i], temperatureB[i] <= 40`


---
## 解题分析及思路：

### 方法：模拟

**思路：**

直接模拟，遍历数组，判断当前元素和下一个元素是否满足条件，满足则计数器加1，否则计数器清零

```go
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	var count int
	var max int
	for i := 0; i < len(temperatureA)-1; i++ {
		if isSame(temperatureA[i+1]-temperatureA[i], temperatureB[i+1]-temperatureB[i]) {
			count++
		} else {
			count = 0
		}
		max = maxFunc(max, count)
	}
	return max
}

func isSame(a, b int) bool {
	if a > 0 && b > 0 {
		return true
	}
	if a < 0 && b < 0 {
		return true
	}
	if a == 0 && b == 0 {
		return true
	}
	return false
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

**复杂度：**

- 时间复杂度：O(N)
- 空间复杂度：O(1)

**执行结果：**

- 执行耗时:0 ms,击败了100.00% 的Go用户
- 内存消耗:3.3 MB,击败了100.00% 的Go用户
