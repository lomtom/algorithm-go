
**题目难度：**[中等](https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/)

**题目描述：**

> 如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
> 
> - n >= 3
> - 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
> 
> 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
> 
> （回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）


**测试用例：**

> 示例 1:
>
> 输入: arr = [1,2,3,4,5,6,7,8]
> 
> 输出: 5
> 
> 解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。

 
> 示例 2：
> 
> 输入: arr = [1,3,7,11,12,14,18]
> 
> 输出: 3
> 
> 解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。


**限制及提示：**
> - 3 <= arr.length <= 1000
> - 1 <= arr[i] < arr[i + 1] <= 10^9

---
**解题分析及思路：**

1、dfs

首先这道题给我的第一直觉就是它能够使用dfs进行解答，事实也是如此。

对于数列序列 X_1, X_2, ..., X_n，满足X_i + X_{i+1} = X_{i+2}，那么只需要每次将前两个满足条件的数值缓存下来，访问下一个元素，如果满足条件更新两个临时值，结果加一，否则跳过该数值，直到遍历完这个数组

其实是很符合dfs的条件的，无奈运行超时，只能放弃，不过思想还是值得大家借鉴的。[最长的斐波那契子序列的长度](https://github.com/lomtom/algorithm-go/blob/main/leetcode/873最长的斐波那契子序列的长度_test.go)

2、dp + hash

对于长度为n的数列，需要为其构建一个n ^ 2的二维数组dp，保存其`dp[raw][col]`位置满足斐波那契序列的组数。

对于`dp[raw][col]`，满足 `arr[raw]` + `arr[col]` = `num`，并且`num`确实是存在与`arr`数组中，其下标为`index`，那么`dp[raw][col]` = `dp[col][index]` + 1，代表存在一组这样的数满足条件。

另一个问题就来了，就算找到数值`num`，那么如何通过`num`判断它是否存在与`arr`中并且找到其所对应下标呢？

可以采用`map`来存`arr`中每一个元素和该元素对应的下标，元素`num`作为key，其下标`index`作为value，那么问题就解决了。

还需要注意两个地方：

- 因为是严格递增的数组，所以内部循环无需遍历n次，到外部循环的所在的下标停止即可。
- 因为设置了`dp[raw][col]` 存放的是满足斐波那契序列的组数，然而题目是返回满足斐波那契序列的元素个数，所以元素个数会比组数多2，在返回结果时加2再返回即可。并且最终结果小于3是无法组成满足斐波那契序列的，返回0即可。

**代码分析：**

1、dfs

首先看超时的dfs：
```go
var result int
var dfs func(arr []int, start, l, num1, num2, num int)
dfs = func(arr []int, start, l, num1, num2, num int) {
    if result < num && num > 2  {
        result = num
    }
    if start >= l {
        return
    }
    for index := start; index < l; index++ {
        if num1+num2 == arr[index] || num1 == 0 || num2 == 0 {
            num++
            dfs(arr, index+1, l, num2, arr[index], num)
            num--
        }
    }
}
```
因为起始状态，`num1`和`num2`并没有存任何的数，这里只需做`num1 == 0 || num2 == 0`条件的额外判断即可。

2、dp + hash

首先定义`map`以及初始化`dp`数组：
```go
l := len(arr)
m := make(map[int]int)
dp := make([][]int, l)
for index := 0; index < l; index++ {
    m[arr[index]] = index
    dp[index] = make([]int, l)
}
```

dp的核心所在：

从二维数组的右下角开始遍历，如果满足条件，将其结果加1并且保存到最终结果中。
```go
for raw := l - 1; raw >= 0; raw-- {
    for col := l - 1; col > raw; col-- {
        if index, ok := m[arr[col]+arr[raw]]; ok {
            dp[raw][col] = dp[col][index] + 1
            ans = max(ans, dp[raw][col])
        }
    }
}
```

最后就是结果的额外处理：
```go
if ans == 0 {
    return
}
return ans + 2
```


最后代码：[最长的斐波那契子序列的长度](https://github.com/lomtom/algorithm-go/blob/main/leetcode/873最长的斐波那契子序列的长度_test.go)

**复杂度：**
> 时间复杂度：O(n ^ 2)，动态规划所需时间
> 
> 空间复杂度：O(n ^ 2)，动态规划数组所需空间

**执行结果：**
> 执行用时： 192 ms , 在所有 Go 提交中击败了 55.84% 的用户
> 
> 内存消耗： 17.3 MB , 在所有 Go 提交中击败了 63.64% 的用户