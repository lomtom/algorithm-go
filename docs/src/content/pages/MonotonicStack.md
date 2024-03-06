---
title: "单调栈"
slug: "monotonicStack"
---
单调栈顾名思义就是满足单调性的栈结构，单调指的是栈内的值要么单调递增，要么单调递减。

那么在应用的时候怎么去维护单调栈的单调性呢？

对于单调递增栈来说，如果要满足栈内元素是递增的，在下一个元素`num`入栈时，需要将栈内比元素`num`大的元素都弹出，然后再将`num`入栈。

因为栈内元素都是单调递增的，所以只需要将栈顶元素与`num`进行比较即可，大于`num`则弹出。

### Go中的栈

那么在Go中怎么来模拟单调栈呢？

众所周知在Go中没有栈这类高级数据结构，那么出栈、入栈都需要自己维护。

首先可以定义一个切片，以切片`stack`来实现栈

```go
stack := make([]int, 0)
```

栈的入栈操作对于切片来说就是切片追加操作，因为在切片末尾加一个元素就是类似于栈顶添加一个元素，即：

```go
stack = append(stack, 1)
```

栈的出栈操作对于切片来说就是将末尾的元素去除，也就是说将切片的长度减一即可，即：

```go
stack = stack[:len(stack) - 1]
```

## 单调栈入栈操作

知道了如何在Go中使用栈，那么单调栈的操作又是怎么样的呢？

例如对于一个`3 1 4 5 2 7`的数组，将数组的元素入栈：

依次判断栈顶元素是否比当前大，大于的话，就执行出栈操作。

- 栈内无元素，3直接入栈；
- 判断栈顶元素3是否比1大，结果为true，将3出栈；
- 栈内无元素，1直接入栈；
- 判断栈顶元素1是否比4大，结果为false，将4入栈；
- 判断栈顶元素4是否比5大，结果为false，将5入栈；
- 判断栈顶元素5是否比2大，结果为true，将5出栈；
- 判断栈顶元素4是否比2大，结果为true，将4出栈；
- 判断栈顶元素1是否比2大，结果为false，将2入栈；
- 判断栈顶元素7是否比2大，结果为false，将7入栈；
- 此时栈内元素为`1,2,7`

![](/img/datastruct/MonotonicStack/stackpush.gif)

对应的代码：

```go
nums := []int{3, 1, 4, 5, 2, 7}
stack := make([]int, 0)
for _, num := range nums {
    for len(stack) > 0 && stack[len(stack)-1] > num {
        // 出栈操作
        stack = stack[:len(stack)-1]
    }
    // 入栈操作
    stack = append(stack, num)
}
// [1 2 7]
fmt.Println(stack)
```

## 模版


其实以上的代码就可以作为一个模版，在实际应用的时候稍微修改一下即可。

```go
循环 数组
    循环 栈不为空 && 栈顶元素大于当前元素
        出栈
    入栈
```

## 使用场景
单调栈可以在时间复杂度为O(n)的情况下，求解出某个元素左边或者右边第一个比它大或者小的元素。


## 例题

[503、下一个更大元素II](leetcode/next-greater-element-ii)

使用单调栈保存元素下标。

每当遍历一个新的数组时，查看栈顶元素对应`nums`中的数值的是否小于当前元素值

- 若小于，将栈顶元素所对应的下标`stack[len(stack)-1]`弹出，并且更新结果集`ans[stack[len(stack)-1]]`的值
- 否则，将当前元素压入栈中。

为了更新全部的结果，需要遍历两次数组，因为考虑到一些额外的情况，例如`[1,2,1]`，如果只遍历一次，那么只有第一个元素更新了相应的结果。而此时，栈中有2,1两个元素。其中1所对应的结果并不知道。
```go
func nextGreaterElements(nums []int) []int {
	l := len(nums)
	stack := make([]int, 0)
	ans := make([]int, l)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < 2*l; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%l] {
			ans[stack[len(stack)-1]] = nums[i%l]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%l)
	}
	return ans
}
```



## 应用

### 中等题

* [503、下一个更大元素II](leetcode/next-greater-element-ii)
* [654、最大二叉树](leetcode/maximum-binary-tree)


