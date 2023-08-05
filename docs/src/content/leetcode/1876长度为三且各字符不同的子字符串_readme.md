---
title: 1876长度为三且各字符不同的子字符串
categories:
  - 简单
tags:
  - 滑动窗口
number: 1876
---

**题目难度：**[长度为三且各字符不同的子字符串](https://leetcode.cn/problems/substrings-of-size-three-with-distinct-characters/) 简单

**题目描述：**

> 如果一个字符串不含有任何重复字符，我们称这个字符串为 好 字符串。
>
> 给你一个字符串 s ，请你返回 s 中长度为 3 的 好子字符串 的数量。
>
> 注意，如果相同的好子字符串出现多次，每一次都应该被记入答案之中。
>
> 子字符串 是一个字符串中连续的字符序列。

**测试用例：**

> 示例 1:
>
> 输入：s = "xyzzaz"
> 
> 输出：1
> 
> 解释：总共有 4 个长度为 3 的子字符串："xyz"，"yzz"，"zza" 和 "zaz" 。
> 
> 唯一的长度为 3 的好子字符串是 "xyz" 。
 
> 示例 2：
>
> 
> 输入：s = "aababcabc"
> 
> 输出：4
> 
> 解释：总共有 7 个长度为 3 的子字符串："aab"，"aba"，"bab"，"abc"，"bca"，"cab" 和 "abc" 。
> 
> 好子字符串包括 "abc"，"bca"，"cab" 和 "abc" 。

**限制及提示：**
> - 1 <= s.length <= 100
> - s 只包含小写英文字母。

---
**解题分析及思路：**

本题 可以采用 [滑动窗口](/window)，只需要维护一个大小为3的窗口，比较窗口内的字符各不相同即可。


**代码分析：**

首先准备大小为3的窗口，即i ,i + 1,i + 2为窗口，并且对其值进行比较，如果不相等，符合条件的结果加一。
```go
func countGoodSubstrings(s string) (ans int) {
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2] {
			ans++
		}
	}
	return ans
}
```
需注意终止条件需要到`len(s)-3`，否则将会越界。


[源代码](https://github.com/lomtom/algorithm-go/blob/main/leetcode/1876长度为三且各字符不同的子字符串_test.go)

**复杂度：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)

**执行结果：**

- 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
- 内存消耗：1.8 MB, 在所有 Go 提交中击败了96.83%的用户
