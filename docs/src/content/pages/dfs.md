---
title: "深度优先搜索"
slug: "dfs"
---
DFS（深度优先搜索）

深度优先搜索 算法（英语：Depth-First-Search，DFS）是一种用于遍历或搜索树或图的算法。
这个算法会 尽可能深 的搜索树的分支。
当结点 v 的所在边都己被探寻过，搜索将 回溯 到发现结点 v 的那条边的起始结点。
这一过程一直进行到已发现从源结点可达的所有结点为止。
如果还存在未被发现的结点，则选择其中一个作为源结点并重复以上过程，整个进程反复进行直到所有结点都被访问为止。


## 应用

### 简单题
* [590、N叉树的后序遍历](../leetcode/n-ary-tree-postorder-traversal)

### 中等题
* [二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)
* [655、输出二叉树](../leetcode/print-binary-tree)
