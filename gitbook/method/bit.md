# 二叉索引树（Binary Indexed Tree）

## 简介

二叉索引树（Binary Indexed Tree）又称树状数组，其发明者又称为Fenwick树。

如果说是二叉索引树，我觉得称之为二进制索引树更为合适，至于为什么，在后面你就能够体会到他的奇妙之处。

最早由Peter M. Fenwick于1994年以《A New Data Structure for Cumulative Frequency Tables》为题发表在SOFTWARE PRACTICE AND EXPERIENCE。
其初衷是解决数据压缩里的累积频率（Cumulative Frequency）的计算问题，现多用于高效计算数列的前缀和、区间和。

所以一般在算法题中求前缀和、前缀积、前缀最大最小的场景就非常适合使用二叉索引树来进行解答。


**值得注意的是能用树状数组解决的问题都可以用其他的解法（例如线段树）来进行解答，反之则不成立。但是这仍不妨碍我们来学习树状数组，因为线段树相对树状数组来说更加复杂，简单的题型使用树状数组往往更合适，主要是学习其中的方法**



某个节点的上一个父节点：i - lowbit(i)

某个节点的父节点：i + lowbit(i)


> 细节处理：需要处理边界：
> 1. 数组长度不变，输出处理边界
> 2. 数组长度加一





## 例题
1. [区域和检索 - 数组不可变](https://leetcode-cn.com/problems/range-sum-query-immutable/)
2. [区域和检索 - 数组可修改](https://leetcode-cn.com/problems/range-sum-query-mutable/)
3. [计算右侧小于当前元素的个数](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/) (选做)
4. [剑指 Offer 51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

## 题解
1、[区域和检索 - 数组不可变](https://leetcode-cn.com/problems/range-sum-query-immutable/)
    这一题使用树状数组并不能体现树状数组的优势，但是那这道题来练一下手也是不错的选择。（前使用缀和数组比较合适）
2、[区域和检索 - 数组可修改](https://leetcode-cn.com/problems/range-sum-query-mutable/)
    是树状数组的经典题目之一，绝佳的练手题型。





[2021.11.23 - ]