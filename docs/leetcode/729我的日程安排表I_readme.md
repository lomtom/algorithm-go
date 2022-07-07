 729. 我的日程安排表 I

---
**题目难度：**[中等](https://leetcode.cn/problems/my-calendar-i/)

**题目描述：**

> 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
>
> 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
> 
> 日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end 。
>
> 实现 MyCalendar 类：
> 
> - MyCalendar() 初始化日历对象。
> 
> - boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false 并且不要将该日程安排添加到日历中。


**测试用例：**

> 示例 ：
>
> 输入：
> 
> ["MyCalendar", "book", "book", "book"]
>
> [[], [10, 20], [15, 25], [20, 30]]
>
> 输出：
>
> [null, true, false, true]
>
> 解释：
>
> MyCalendar myCalendar = new MyCalendar();
>
> myCalendar.book(10, 20); // return True
>
> myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15 已经被另一个日程安排预订了。
>
> myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于 20 ，且不包含时间 20 。


**限制：**
> 0 <= start < end <= 10^9
> 
> 每个测试用例，调用 book 方法的次数最多不超过 1000 次。

---

**解题分析及思路：**

不难看出，其实这是一道区间的判断问题，我们只需要判断在不在某些特定的区间即可。


那么可以将这个区间用一个数组进行表示，那么数组的长度是 10 ^ 9，但是实际上这那个不可取，空间上超过了限制，并且时间上复杂度也会过高。所以直接pass该做法。

那么怎么基于该做法做一个优化呢？

我们可以将 [start, end) 的两端保存起来，可以采用二维数组，只要能够保存两个数就行了。

然后需要着重关注的是**边界问题**。

例如，对于已经存在 10 - 20 的区间，再push一个区间，可能会有以下情况：

```
   10 - 20
1. 20 - 30 true
2. 11 - 15 false
3. 10 - 15 false
4. 15 - 20 false
5. 8  - 20 false
6. 8  - 22 false
7. 8  - 10 true
8. 8  - 15 false
9. 15 - 22 false
```

那么可以将结果为false的情况归纳为三种：（当然，也可以判断为true的情况）
1. start 处于 10 - 20 之间
2. end 处于 10 - 20 之间
3. start 小于 10 ，end 大于 20（容易忽略）

然后根据结果保存值，不符合上述情况，将start、end添加到切片末尾。

下一个start、end到来的时候，遍历切片再次判断即可

**代码分析：**
定义并初始化一个切片：
```go
type MyCalendar struct {
    time [][2]int
}

func Constructor() MyCalendar {
    return MyCalendar{
        time: make([][2]int, 0),
    }
}
```

每次进行book的时候对false的情况进行判断，如果部分条件，直接返回false。
```go
for index := range this.time {
    if this.time[index][0] <= start && start < this.time[index][1] || this.time[index][0] < end && end < this.time[index][1] || start <= this.time[index][0] && this.time[index][1] <= end {
        return false
    }
}
```
符合条件，将start、end以数组的形式append到切片，并且返回为true。
```go
this.time = append(this.time, [2]int{start, end})
return true
```


最后代码：[我的日程安排表I](https://github.com/lomtom/algorithm-go/blob/main/leetcode/729我的日程安排表I_test.go)

**复杂度：**
> 时间复杂度：O(n^2)，遍历切片所需时间
> 空间复杂度：O(n)，切片所需空间

**执行结果：**
> 执行用时： 80 ms , 在所有 Go 提交中击败了 49.10% 的用户
> 
> 内存消耗： 7 MB , 在所有 Go 提交中击败了 74.25% 的用户