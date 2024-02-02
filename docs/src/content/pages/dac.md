---
title: "分治法"
slug: "dac"
---
## 简介

分治法可以通俗的解释为：把一片领土分解，分解为若干块小部分，然后一块块地占领征服，被分解的可以是不同的政治派别或是其他什么，然后让他们彼此异化。

分治法的精髓：

1. 分--将问题分解为规模更小的子问题；
2. 治--将这些规模更小的子问题逐个击破；
3. 合--将已解决的子问题合并，最终得出“母”问题的解；

分治法所能解决的问题一般具有以下几个特征：

1. 该问题的规模缩小到一定的程度就可以容易地解决
2. 该问题可以分解为若干个规模较小的相同问题，即该问题具有最优子结构性质。
3. 利用该问题分解出的子问题的解可以合并为该问题的解；
4. 该问题所分解出的各个子问题是相互独立的，即子问题之间不包含公共的子问题。

## 分治法应用:

### 简单：
- [104、二叉树的最大深度](leetcode/maximum-depth-of-binary-tree) ⭐
- [110、平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/) ⭐

### 中等
* [241、为运算表达式设计优先级](https://leetcode.cn/problems/different-ways-to-add-parentheses/)
* [558、四叉树交集](leetcode/logical-or-of-two-binary-grids-represented-as-quad-trees)
* [814、二叉树剪枝](leetcode/binary-tree-pruning)
* [654、最大二叉树](leetcode/maximum-binary-tree)
