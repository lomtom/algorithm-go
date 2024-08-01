---
title: 心算挑战
categories:
  - 简单
  - LCP
tags: 
  - 贪心
  - 数组
  - 排序
slug: uOAnQW
number: LCP40
---

## 题目描述：

「力扣挑战赛」心算项目的挑战比赛中，要求选手从 `N` 张卡牌中选出 `cnt` 张卡牌，若这 `cnt` 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 `cnt` 张卡牌数字总和。
给定数组 `cards` 和 `cnt`，其中 `cards[i]` 表示第 `i` 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。

示例 1：
```
输入：`cards = [1,2,8,9], cnt = 3`

输出：`18`

解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。

```
示例 2：
```
输入：`cards = [3,3,1], cnt = 1`

输出：`0`

解释：不存在获取有效得分的卡牌方案。

```
**提示：**
- `1 <= cnt <= cards.length <= 10^5`
- `1 <= cards[i] <= 1000`




---
## 解题分析及思路：

### 方法：排序 + 贪心

**思路：**
将 cards 从大到小排序后，先贪心的将后 cnt 个数字加起来，若此时 sum 为偶数，直接返回即可。

若此时答案为奇数，有两种方案:
- 在数组前面找到一个最大的奇数与后 cnt 个数中最小的偶数进行替换；
- 在数组前面找到一个最大的偶数与后 cnt 个数中最小的奇数进行替换。

两种方案选最大值即可。

```go
import "sort"

func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	var result = 0
	var odd, even = -1, -1
	for index := 0; index < cnt; index++ {
		result += cards[index]
		if cards[index]%2 == 0 {
			even = index
		} else {
			odd = index
		}
	}
	if result%2 == 0 {
		return result
	}
	for index := cnt; index < len(cards); index++ {
		if cards[index]%2 == 1 && even != -1 {
			result += cards[index] - cards[even]
			return result
		} else if cards[index]%2 == 0 && odd != -1 {
			result += cards[index] - cards[odd]
			return result
		}
	}
	return 0
}
```

**复杂度：**

- 时间复杂度：O(n)
- 空间复杂度：O(1)

**执行结果：**

- 执行耗时:366 ms,击败了42.86% 的Go用户
- 内存消耗:8.4 MB,击败了14.29% 的Go用户
