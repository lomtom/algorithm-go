---
title: Pow(x, n)
categories:
  - 中等
tags:
  - 数学
  - 快速幂
slug: powx-n
number: 50
---

## 题目描述：

实现 [pow(x, n)](https://cplusplus.com/reference/valarray/pow/) ，即计算 x 的整数 n 次幂函数（即，xn ）。

**示例 1：**
```
输入：x = 2.00000, n = 10
输出：1024.00000
```

**示例 2：**
```
输入：x = 2.10000, n = 3
输出：9.26100
```

**示例 3：**
```
输入：x = 2.00000, n = -2
输出：0.25000
解释：2^-2 = 1/2^2 = 1/4 = 0.25
```


**提示：**
- -100.0 < x < 100.0
- -2<sup>31</sup> <= n <= 2<sup>31</sup>-1
- n 是一个整数
- 要么 x 不为零，要么 n > 0 。
- -10<sup>4</sup> <= xn <= 10<sup>4</sup>

---
## 解题分析及思路：



### 方法：快速幂 + 迭代

快速幂算法是通过二进制的方式来计算幂的，例如计算 x<sup>11</sup>，可以通过 x<sup>8</sup> * x<sup>2</sup> * x<sup>1</sup> 来计算。

具体步骤如下：

1. 初始化结果为 1，即 res = 1
2. 从低位到高位遍历 n 的二进制位，如果当前位为 1，则将结果乘上 x 的对应幂次
3. x 的幂次每次翻倍，即 x = x * x
4. n 的二进制位每次右移一位，即 n = n >> 1
5. 重复步骤 2-4 直到 n 为 0
6. 返回结果

如果 n 为负数，则将 x 的 n 次幂转换为 1/(x 的 -n 次幂)。

```go
func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, -n)
}

func pow(x float64, n int) float64 {
	res := 1.0
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
	}
	return res
}
```

**复杂度：**

- 时间复杂度：O(logN) 
- 空间复杂度：O(1)

**执行结果：**

- 执行耗时:2 ms,击败了21.50% 的Go用户
- 内存消耗:2.1 MB,击败了8.43% 的Go用户
